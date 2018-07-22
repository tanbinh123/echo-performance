# echo performance
echo performance verification repository.

# Docker image
Docker pull
```
$ docker pull tomohito/echo-performance
```

Set Environment variable.
```
DB_SOURCE={database source}
```

Docker run.
```
$ docker run -p 8080:8080 tomohito/echo-performance
```

# Build by myself
Git clone.
```
$ git clone git@github.com:tomoyane/echo-performance.git
```

Docker build.
```
$ docker build -t {username}/echo-performance:latest .
```

Set Environment variable.
```
DB_SOURCE={database source}
```

Docker run.
```
$ docker run -p 8080:8080 {username}/echo-performance
```

# Performance verification sql
```
CREATE TABLE items (
  id INT(11) NOT NULL AUTO_INCREMENT,
  name VARCHAR(128) NOT NULL,
  category VARCHAR(128) NOT NULL,
  created_at DATETIME(6),
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```
