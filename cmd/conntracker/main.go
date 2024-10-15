package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/epfl-si/mariadb-conntracker/internal/conntracker"
)

var version = "dev"

func main() {

	versionFlag := flag.Bool("version", false, "Print version information and exit")
	flag.Parse()

	if *versionFlag || (len(os.Args) > 1 && os.Args[1] == "version") {
		fmt.Printf("conntracker version %s\n", version)
		os.Exit(0)
	}

	cfg, err := conntracker.InitConfig()
	if err != nil {
		log.Fatalf("Failed to initialize config: %v", err)
	}

	var programLevel = new(slog.LevelVar)
	h := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})
	slog.SetDefault(slog.New(h))
	programLevel.Set(cfg.LogLevel)

	if err := run(cfg); err != nil {
		log.Fatal(err)
	}
}

func run(cfg conntracker.Config) error {

	db, err := conntracker.OpenOrCreateDB(cfg)
	if err != nil {
		slog.Error("failed to open or create database.", "error_msg", err)
		return err
	}
	defer db.Close()

	lastProcessedTime, err := conntracker.GetLastProcessedTime(cfg, db)
	if err != nil {
		slog.Error("failed to get last processed time", "error_msg", err)
		return err
	}

	filePaths, newLastProcessedTime, err := conntracker.FilterAndSortNewFiles(cfg, lastProcessedTime)
	if err != nil {
		slog.Error("error getting files to process", "error_msg", err)
		return err
	}

	// Round to the nearest second (filesystem returns 2024-10-15 13:43:57.109984656 +0200 CEST)
	roundedNewLastProcessedTime := newLastProcessedTime.Round(time.Second)

	if lastProcessedTime.Compare(roundedNewLastProcessedTime) >= 0 {
		slog.Info("no connections found since last processed time", "lastProcessedTime", lastProcessedTime)
		return nil
	}

	accounts, err := conntracker.ProcessFilesParallel(cfg, filePaths)
	if err != nil {
		slog.Error("error processing files", "error_msg", err)
		return err
	}
	slog.Debug("debug accounts found accross all files", "accounts_parsed", len(accounts))

	inserted, updated, err := conntracker.InsertOrUpdateAccounts(cfg, db, accounts)
	if err != nil {
		slog.Error("failed to insert or update accounts", "error_msg", err)
		return err
	}
	slog.Debug("debug affected accounts count", "inserted", inserted, "updated", updated)

	if err := conntracker.UpdateLastProcessedTime(cfg, db, newLastProcessedTime); err != nil {
		slog.Error("error updating the date of the last parsing.", "error_msg", err)
		return err
	}

	return nil
}
