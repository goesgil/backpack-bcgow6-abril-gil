package transactions

import (
	"encoding/json"
	"testing"

	"github.com/goesgil/backpack-bcgow6-abril-gil/testing/clase_3/pkg/store"
	"github.com/stretchr/testify/assert"
)

/*Ejercicio 1 - Service/Repo/Db Update()

Diseñar un test que pruebe en la capa service, el método o función Update(). Para lograrlo se deberá:
	1. Crear un mock de Storage, dicho mock debe contener en su data un producto con las especificaciones que desee.
	2. El método Read del Mock, debe contener una lógica que permita comprobar que dicho método fue invocado.
	3. Para dar el test como OK debe validarse que al invocar el método del Service Update(),  retorne el producto con
	mismo Id y los datos actualizados. Validar también que  Read() del Store haya sido ejecutado durante el test.
*/

func TestUpdate(t *testing.T) {
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

	mockUpdatedTrx := &Transaction{
		Id:          1,
		CodeTrxs:    "code-updated",
		Amount:      400,
		Transmitter: "Persona3",
		Coin:        "ETH",
	}

	repository := NewRepository(&stubStore)
	service := NewService(repository)
	trx, err := service.Update(1, 400, "code-updated", "ETH", "Persona3")
	assert.Nil(t, err)
	assert.Equal(t, mockUpdatedTrx, trx)
}

/*Ejercicio 2 - Service/Repo/Db Delete()
Diseñar un test que pruebe en la capa service, el método o función Delete(). Se debe probar la
correcta eliminación de un producto, y el error cuando el producto no existe. Para lograrlo puede:
	1. Crear un mock de Storage, dicho mock debe contener en su data un producto con las especificaciones que desee.
	2. Ejecutar el test con dos id’s de producto distintos, siendo uno de ellos un id inexistente en el Mock de Storage.
	3. Para dar el test como OK debe validarse que efectivamente el producto borrado ya no exista en Storage luego del Delete().
	También que cuando se intenta borrar un producto  inexistente, se debe obtener el error correspondiente.
*/

func TestDelete(t *testing.T) {
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

	repository := NewRepository(&stubStore)
	service := NewService(repository)
	err := service.Delete(1)

	assert.Nil(t, err)
}
