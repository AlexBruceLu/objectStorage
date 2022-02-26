export RABBIT_SERVER=amqp://test:test@172.24.218.136:5672

dev:
  @sh scripts/start.sh
  @echo started.

init:
  @sh scripts/init.sh