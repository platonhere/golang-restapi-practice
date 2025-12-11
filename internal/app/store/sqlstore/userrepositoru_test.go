package sqlstore_test

import (
	"http-rest-api/internal/app/model"
	"http-rest-api/internal/app/store"
	"http-rest-api/internal/app/store/sqlstore"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestStore создаёт тестовую базу данных и проверяет успешное открытие соединения.
// Возвращает объект Store и функцию teardown, которая очищает указанные таблицы
// и связанные с н ими данные после теста.
// В тесте вызывается метод Create репозитория UserRepository и проверяется,
// что ошибка отсутствует, а возвращённый пользователь не nil.

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)
	s.User().Create(u)
	u, err := s.User().Find(u.ID)

	assert.NoError(t, err)
	assert.NotNil(t, u)

}

// TestStore открывает соединение с тестовой БД, defer teardown("users") очищает таблицу после теста.
// Тест проверяет метод FindByEmail интерфейса UserRepository.
// Сначала проверяется, что поиск пользователя в пустой таблице возвращает ошибку.
// Затем создаётся пользователь с указанным email, и повторный поиск должен пройти успешно:
// ошибки нет, а возвращённый пользователь не nil.
func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)
	_, err := s.User().FindByEmail(u.Email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.User().Create(u)
	u, err = s.User().FindByEmail(u.Email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
