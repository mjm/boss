project "test" {
  task "task1" {
    cmd = "while true; do echo hello; sleep 2; done"
  }

  task "task2" {
    cmd = "log stream"
  }
}
