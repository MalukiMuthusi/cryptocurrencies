package fauna

import (
	"context"

	"github.com/MalukiMuthusi/cryptocurrencies/internal/logger"
	"github.com/MalukiMuthusi/cryptocurrencies/internal/models"
	"github.com/fauna/faunadb-go/v4/faunadb"
	"github.com/spf13/viper"
)

// Fauna database
type Fauna struct {
	Client *faunadb.FaunaClient
}

// New creates a new client connection to the Fauna database
func New() (*Fauna, error) {

	client := faunadb.NewFaunaClient(viper.GetString("FAUNA_SECRET"))

	return &Fauna{Client: client}, nil
}

// List cryptocurrencies
func (f Fauna) List(c context.Context) ([]*models.ListRes, error) {

	res, err := f.Client.Query(faunadb.Map(faunadb.Paginate(faunadb.Documents(faunadb.Collection("cryptocurrencies"))), faunadb.Lambda("ref", faunadb.Get(faunadb.Var("ref")))))

	if err != nil {
		logger.Log.Info(err)
		return nil, err
	}

	var cryptocurrencies []*models.ListRes

	logger.Log.Info(res)

	if err := res.At(faunadb.ObjKey("data")).Get(&cryptocurrencies); err != nil {
		logger.Log.Info(err)
		return nil, err
	}

	return cryptocurrencies, nil
}
