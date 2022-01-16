package cmd

import (
	"fmt"
	"log"

	"github.com/mgutz/ansi"
	"github.com/ostamand/simpurl/internal/config"
	"github.com/ostamand/simpurl/internal/store"
	"github.com/ostamand/simpurl/internal/store/mysql"
	"github.com/spf13/cobra"
)

var port string

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed database with default development data",
	Long: `Simplify development by seeding automatically the database with develoment data.

Do not run this on the production environment.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		params := config.Get(cfgFile)

		if port != "" {
			params.Db.Port = port
		}

		s := mysql.InitializeSQL(&params.Db)

		s.User.Save(
			&store.UserModel{
				Username: "admin",
				Password: "admin",
				Admin: true,
			},
		)

		s.User.Save(
			&store.UserModel{
				Username: "user",
				Password: "user",
				Admin: false,
			},
		)

		u, _ := s.User.GetByUsername("admin")

		// save 10 links for admin
		for i :=1; i <=10; i++ {
			s.Link.Save(
				&store.LinkModel{
					UserID: u.ID,
					Symbol: fmt.Sprintf("symbol_%d", i),
					URL: fmt.Sprintf("https://example_%d.com", i),
					Description: fmt.Sprintf("description_%d", i),
					Note: fmt.Sprintf("note_%d", i), 
				},
			)
		}

		phosphorize := ansi.ColorFunc("green+h:black")
		log.Println(phosphorize("Database seeded"))
	},
}

func init() {
	dbCmd.AddCommand(seedCmd)
	seedCmd.Flags().StringVarP(&port, "port", "p", "", "Overwrite MySQL port in config")
}
