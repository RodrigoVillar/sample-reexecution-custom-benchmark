window.BENCHMARK_DATA = {
  "lastUpdate": 1764857698575,
  "repoUrl": "https://github.com/RodrigoVillar/sample-reexecution-custom-benchmark",
  "entries": {
    "Benchmark": [
      {
        "commit": {
          "author": {
            "email": "rodrigo.villar@avalabs.org",
            "name": "Rodrigo Villar",
            "username": "RodrigoVillar"
          },
          "committer": {
            "email": "rodrigo.villar@avalabs.org",
            "name": "Rodrigo Villar",
            "username": "RodrigoVillar"
          },
          "distinct": true,
          "id": "02bc12ec093bc51203a06e1169443e8acc37d13e",
          "message": "chore: use correct go version + invokation",
          "timestamp": "2025-12-04T09:09:33-05:00",
          "tree_id": "3abd3a6320e7144a435faad8d998c124a0770289",
          "url": "https://github.com/RodrigoVillar/sample-reexecution-custom-benchmark/commit/02bc12ec093bc51203a06e1169443e8acc37d13e"
        },
        "date": 1764857698323,
        "tool": "customBiggerIsBetter",
        "benches": [
          {
            "name": "BenchmarkReexecuteRange/[1,50000]-Config-default-Runner-dev",
            "value": 1.1,
            "unit": "block_accept_ms/ggas"
          },
          {
            "name": "BenchmarkReexecuteRange/[1,50000]-Config-default-Runner-dev",
            "value": 2.2,
            "unit": "block_parse_ms/ggas"
          },
          {
            "name": "BenchmarkReexecuteRange/[1,50000]-Config-default-Runner-dev",
            "value": 3.3,
            "unit": "block_verify_ms/ggas"
          },
          {
            "name": "BenchmarkReexecuteRange/[1,50000]-Config-default-Runner-dev",
            "value": 4,
            "unit": "mgas/s"
          }
        ]
      }
    ]
  }
}