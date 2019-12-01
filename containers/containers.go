package containers

import (
	"fmt"
	"log"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mintutu/restful-golang/config"
	"github.com/mintutu/restful-golang/driver"
	ph "github.com/mintutu/restful-golang/handler/http"
	"github.com/mintutu/restful-golang/repository"
	"github.com/mintutu/restful-golang/repository/post"
	"github.com/spf13/viper"
	"go.uber.org/dig"
)

var Container *dig.Container

func LoadConfiguration() *config.Configuration {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	var configuration config.Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	return &configuration
}

func ConnectDatabase(configuration *config.Configuration) *driver.DB {

	connection, err := driver.ConnectSQL(
		configuration.Database.Host,
		configuration.Database.Port,
		configuration.Database.User,
		configuration.Database.Pass,
		configuration.Database.Name)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	return connection
}

func SQLRepoManager(connection *driver.DB) repository.PostRepo {
	return post.NewMySQLPostRepo(connection.SQL)
}

func RestRepoManager() repository.PostRepo {
	return post.NewRestPostRepo()
}

func HandlerManager(repo repository.PostRepo) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	pHandler := ph.NewPostHandler(repo)
	cHandler := ph.NewCommentHandler()
	r.Get("/comments", cHandler.Fetch)
	r.Get("/posts", pHandler.Fetch)
	r.Post("/posts", pHandler.Create)
	r.Mount("/debug", middleware.Profiler())
	return r
}

func BuildContainer() *dig.Container {
	Container := dig.New()
	Container.Provide(LoadConfiguration)
	Container.Provide(ConnectDatabase)
	// Container.Provide(RestRepoManager)
	Container.Provide(SQLRepoManager)
	Container.Provide(HandlerManager)
	return Container
}
