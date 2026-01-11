import multiprocessing

bind = "0.0.0.0:8000"
backlog = 2048

workers = multiprocessing.cpu_count() * 2 + 1
worker_class = "sync"
worker_connections = 1000
timeout = 30
graceful_timeout = 30
keepalive = 2

accesslog = "-"
errorlog = "-"
loglevel = "info"
capture_output = True

proc_name = "analysis_service"

limit_request_line = 4094
limit_request_fields = 100
limit_request_field_size = 8190

reload = False
