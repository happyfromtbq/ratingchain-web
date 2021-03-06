package repositories

import (
	"sync"

	"github.com/happyfromtbq/ratingchain-web/datamodels"
	"fmt"
)

// Query represents the visitor and action queries.
type Query func(datamodels.User) bool

// UserRepository handles the basic operations of a user entity/model.
// It's an interface in order to be testable, i.e a memory user repository or
// a connected to an sql database.
type UserRepository interface {
	Exec(query Query, action Query, limit int, mode int) (ok bool)

	Select(query Query) (user datamodels.User, found bool)
	SelectMany(query Query, limit int) (results []datamodels.User)

	InsertOrUpdate(user datamodels.User) (updatedUser datamodels.User, err error)
	Delete(query Query, limit int) (deleted bool)
}

// NewUserRepository returns a new user memory-based repository,
// the one and only repository type in our example.
func NewUserRepository(source map[int64]datamodels.User) UserRepository {
	return &userMemoryRepository{source: source}
}

// userMemoryRepository is a "UserRepository"
// which manages the users using the memory data source (map).
type userMemoryRepository struct {
	source map[int64]datamodels.User
	mu     sync.RWMutex
}

const (
	// ReadOnlyMode will RLock(read) the data .
	ReadOnlyMode = iota
	// ReadWriteMode will Lock(read/write) the data.
	ReadWriteMode
)

func (r *userMemoryRepository) Exec(query Query, action Query, actionLimit int, mode int) (ok bool) {
	loops := 0

	if mode == ReadOnlyMode {
		r.mu.RLock()
		defer r.mu.RUnlock()
	} else {
		r.mu.Lock()
		defer r.mu.Unlock()
	}

	for _, user := range r.source {
		ok = query(user)
		if ok {
			if action(user) {
				loops++
				if actionLimit >= loops {
					break // break
				}
			}
		}
	}

	return
}

// Select receives a query function
// which is fired for every single user model inside
// our imaginary data source.
// When that function returns true then it stops the iteration.
//
// It returns the query's return last known boolean value
// and the last known user model
// to help callers to reduce the LOC.
//
// It's actually a simple but very clever prototype function
// I'm using everywhere since I firstly think of it,
// hope you'll find it very useful as well.
func (r *userMemoryRepository) Select(query Query) (user datamodels.User, found bool) {
	found = r.Exec(query, func(m datamodels.User) bool {
		user = m
		return true
	}, 1, ReadOnlyMode)

	// set an empty datamodels.User if not found at all.
	if !found {
		user = datamodels.User{}
	}

	return
}

// SelectMany same as Select but returns one or more datamodels.User as a slice.
// If limit <=0 then it returns everything.
func (r *userMemoryRepository) SelectMany(query Query, limit int) (results []datamodels.User) {
	r.Exec(query, func(m datamodels.User) bool {
		results = append(results, m)
		return true
	}, limit, ReadOnlyMode)

	return
}

// InsertOrUpdate adds or updates a user to the (memory) storage.
//
// Returns the new user and an error if any.
func (r *userMemoryRepository) InsertOrUpdate(user datamodels.User) (datamodels.User, error) {

	InitDb()
	data := map[string]interface{}{
		"username":  user.Username,
		"salt":  "salt",
		"password": user.Password,
	}
	defer conn.Close()
	db := conn.NewDB()
	db2 :=db.Table("t_user").Data(data);
	res, err := InsertGetId(db2);
	//res, err := db.Table("t_user").Data(data).InsertGetId() //int64 to  int
	fmt.Println(res)
	if err != nil {
		fmt.Println(err)
		var user2 datamodels.User
		return user2,nil
	}else{
		user.ID = res
		return user,nil
	}
}

func (r *userMemoryRepository) Delete(query Query, limit int) bool {
	return r.Exec(query, func(m datamodels.User) bool {
		delete(r.source, m.ID)
		return true
	}, limit, ReadWriteMode)
}
