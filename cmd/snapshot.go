package cmd

import (
	"github.com/ironstar-io/tokaido/initialize"
	"github.com/ironstar-io/tokaido/services/docker"
	"github.com/ironstar-io/tokaido/services/snapshots"
	"github.com/ironstar-io/tokaido/utils"
	"github.com/spf13/cobra"
)

var nameFlag string

// SnapshotCmd - `tok snapshot new`
var SnapshotCmd = &cobra.Command{
	Use:   "snapshot [name]",
	Short: "Manage database snapshots saved in .tok/local/snapshots",
	Long:  "Create, restore, database snapshots saved to .tok/local/snapshots. These can be created by Tokaido, or saved SQL dumps from somewhere else.",
	Run: func(cmd *cobra.Command, args []string) {
		initialize.TokConfig("tokaido")
		utils.CheckCmdHard("docker-compose")

		docker.HardCheckTokCompose()

		snapshots.New(args)
	},
}

// SnapshotNewCmd - `tok snapshot new`
var SnapshotNewCmd = &cobra.Command{
	Use:   "new [name]",
	Short: "Creates a new database snapshot",
	Long:  "Creates a new database snapshot and saves it to .tok/local/snapshots with the current UTC date",
	Run: func(cmd *cobra.Command, args []string) {
		initialize.TokConfig("tokaido")
		utils.CheckCmdHard("docker-compose")

		docker.HardCheckTokCompose()

		snapshots.New(args)
	},
}

// SnapshotListCmd - `tok snapshot list`
var SnapshotListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all snapshots with their current ID",
	Long:  "Lists all snapshots with their current ID",
	Run: func(cmd *cobra.Command, args []string) {
		initialize.TokConfig("tokaido")
		utils.CheckCmdHard("docker-compose")

		docker.HardCheckTokCompose()

		snapshots.List()
	},
}

// SnapshotDeleteCmd - `tok snapshot cleanup`
var SnapshotDeleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Deletes all tokaido snapshots from the .tok/local/snapshots directory",
	Long:  "Creates a new database snapshot and saves it to .tok/local/snapshots with the current UTC date",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		initialize.TokConfig("tokaido")
		utils.CheckCmdHard("docker-compose")

		docker.HardCheckTokCompose()

		snapshots.Delete(args)
	},
}

// SnapshotRestoreCmd - `tok snapshot cleanup`
var SnapshotRestoreCmd = &cobra.Command{
	Use:   "restore [id]",
	Short: "Restores the specified snapshot id, or provides a list to choose from",
	Long:  "Restores the specified snapshot id, or provides a list to choose from",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		initialize.TokConfig("tokaido")
		utils.CheckCmdHard("docker-compose")

		docker.HardCheckTokCompose()

		snapshots.Restore(args)
	},
}
