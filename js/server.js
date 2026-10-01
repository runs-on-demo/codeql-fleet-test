const express = require("express");
const { exec } = require("child_process");

const app = express();

app.get("/ping", (req, res) => {
  // Deliberate command injection for the CodeQL test.
  exec("ping -c 1 " + req.query.host, (err, stdout) => {
    res.send(stdout);
  });
});

app.listen(3000);
