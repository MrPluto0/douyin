wrk.method = "GET"
wrk.host = "127.0.0.1"
wrk.port = 8080
wrk.body = ''
wrk.headers["Content-Type"] = "application/json"
wrk.path = '/douyin/feed'