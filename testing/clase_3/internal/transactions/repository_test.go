package transactions

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/goesgil/backpack-bcgow6-abril-gil/testing/clase_3/pkg/store"
	"github.com/stretchr/testify/assert"
)

/*
Ejercicio 1 - Test Unitario GetAll()
Generar un Stub del Store cuya función “Read” retorne dos productos con las especificaciones que deseen. Comprobar que GetAll() retorne la información exactamente igual a la esperada. Para esto:
    1. Dentro de la carpeta /internal/products, crear un archivo repository_test.go con el test diseñado.
*/

func TestGetAll(t *testing.T) {
	trxs := []*Transaction{
		{
			Id:          1,
			CodeTrxs:    "code-1",
			Amount:      500,
			Transmitter: "Persona1",
			Coin:        "BTC",
		},
		{
			Id:          2,
			CodeTrxs:    "code-2",
			Amount:      500.50,
			Transmitter: "Persona1",
			Coin:        "BTC",
		},
	}

	data, _ := json.Marshal(trxs)
	dbMock := store.Mock{
		Data:  data,
		Error: nil,
	}

	stubStore := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}

	r := NewRepository(&stubStore)
	responseGetAllTrxs, err := r.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, trxs, responseGetAllTrxs)
}

func TestGetAllError(t *testing.T) {
	errorExpected := errors.New("error expected")
	dbMock := store.Mock{
		Error: errorExpected,
	}

	stubStore := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}

	r := NewRepository(&stubStore)
	trxs, err := r.GetAll()

	assert.Equal(t, errorExpected, err)
	assert.Nil(t, trxs)
}

/*
Ejercicio 2 - Test Unitario UpdateName()
Diseñar Test de UpdateName, donde se valide que la respuesta retornada sea correcta para la actualización del nombre de un producto específico. Y además se compruebe que efectivamente se usa el método “Read” del Storage para buscar el producto. Para esto:
    1. Crear un mock de Storage, dicho mock debe contener en su data un producto específico cuyo nombre puede ser “Before Update”.
    2. El método Read del Mock, debe contener una lógica que permita comprobar que dicho método fue invocado. Puede ser a través de un boolean como se observó en la clase.
    3. Para dar el test como OK debe validarse que al invocar el método del Repository UpdateName, con el id del producto mockeado y con el nuevo nombre “After Update”, efectivamente haga la actualización. También debe validarse que el método Read haya sido ejecutado durante el test.
*/

func TestUpdateName(t *testing.T) {
	id, codeTrxs := 1, "code-updated"
	trxs := []*Transaction{
		{
			Id:          1,
			CodeTrxs:    "code-1",
			Amount:      500,
			Transmitter: "Persona1",
			Coin:        "BTC",
		},
	}

	data, _ := json.Marshal(trxs)
	mock := store.Mock{Data: data}

	stubStore := store.FileStore{
		FileName: "",
		Mock:     &mock,
	}

	r := NewRepository(&stubStore)
	updatedTrxs, err := r.UpdateName(id, codeTrxs)
	assert.Nil(t, err)

	assert.True(t, mock.ReadInvoked)
	assert.Equal(t, id, updatedTrxs.Id)
	assert.Equal(t, codeTrxs, updatedTrxs.CodeTrxs)
}
