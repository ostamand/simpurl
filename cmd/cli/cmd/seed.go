package cmd

import (
	"log"
	"math/rand"

	"github.com/brianvoe/gofakeit/v6"
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
		for i :=1; i <=20; i++ {
			symbol := ""
			if rand.Intn(100) > 80 {
				symbol = gofakeit.Word()
			}
			s.Link.Save(
				&store.LinkModel{
					UserID: u.ID,
					Symbol: symbol,
					URL: gofakeit.URL(),
					Description: gofakeit.SentenceSimple(),
					Note: gofakeit.Sentence(12), 
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
