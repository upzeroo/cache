# General

Simple cache component. Atm has only Redis implementation, but implementinbg/dealing with/ others is easy - just satisfy the cache.Adapter interface, instantiate and inject it into the cache service - thats it. 

# Example

Lets take alook at this sample/illustrative lines:
```
package main

import (
...
	"github.com/upzeroo/cache/adapters"
	cService "github.com/upzeroo/cache/service"
)

type (
        Something struct{
            ...
	        cache    *cService.CacheService
        }
)

func main() {
	logger := logger.Create().WithFields(logrus.Fields{
        ...
	})

	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file")
	}

	cacheAdapter, err := adapters.Factory("redis", &adapters.DepContainer{
		RedisURL: os.Getenv("REDIS_URL"),
	})
	if err != nil {
		logger.Fatal("redis url not defined")
	}
	
    cacheService := cService.NewCacheService(cacheAdapter, logger)
    ...

    s := &Something{
        cache: cacheService,
    }
    
    s.DoOne
    ...
    s.DoOther
    ...
}

func (s *Something) DoOne() error {
    ...

	res, err := s.cache.Get("some-key-here")
	if err != nil {
        return err     
    }
    ...
}


func (s *Something) DoOther() error {
    ...

    err = s.cache.Set(cacheKey, struct {
        URL    string
        Status string
    }{
        URL:    urlStr,
        Status: "is valid",
    }, 20 * time.Second)

    ...
}

...
/// same for cacheService.Delete

```

# Test

Run:
```
/storage/.../upzeroo/cache >>> godotenv -f .env go test -v ./... -cover                                                                                               ±[●][main]
?       github.com/upzeroo/cache        [no test files]
?       github.com/upzeroo/cache/service        [no test files]
=== RUN   Test_Factory
=== RUN   Test_Factory/redis_-_should_pass
=== RUN   Test_Factory/non_existent_adapter_-_should_err
=== RUN   Test_Factory/empty_adapter_-_should_err
--- PASS: Test_Factory (0.00s)
    --- PASS: Test_Factory/redis_-_should_pass (0.00s)
    --- PASS: Test_Factory/non_existent_adapter_-_should_err (0.00s)
    --- PASS: Test_Factory/empty_adapter_-_should_err (0.00s)
=== RUN   Test_NewRedisAdapter
=== RUN   Test_NewRedisAdapter/redis_-_should_pass
=== RUN   Test_NewRedisAdapter/redis_-_broker_ulr_-_should_err
--- PASS: Test_NewRedisAdapter (0.00s)
    --- PASS: Test_NewRedisAdapter/redis_-_should_pass (0.00s)
    --- PASS: Test_NewRedisAdapter/redis_-_broker_ulr_-_should_err (0.00s)
=== RUN   Test_Get
=== RUN   Test_Get/get_real_key_-_must_pass
=== RUN   Test_Get/get_non_existent_-_must_err
--- PASS: Test_Get (0.00s)
    --- PASS: Test_Get/get_real_key_-_must_pass (0.00s)
    --- PASS: Test_Get/get_non_existent_-_must_err (0.00s)
=== RUN   Test_Set
=== RUN   Test_Set/set_-_must_pass
--- PASS: Test_Set (0.00s)
    --- PASS: Test_Set/set_-_must_pass (0.00s)
=== RUN   Test_Delete
=== RUN   Test_Delete/delete_-_must_pass
--- PASS: Test_Delete (0.00s)
    --- PASS: Test_Delete/delete_-_must_pass (0.00s)
PASS
        github.com/upzeroo/cache/adapters       coverage: 75.0% of statements
ok      github.com/upzeroo/cache/adapters       0.005s  coverage: 75.0% of statements
```
