const express = require("express");
const bodyParser = require("body-parser");
const notifier = require("node-notifier");
const app = express();
const port = process.env.PORT || 9000;
const path = require("path");

app.use(bodyParser.json());

app.get("/health", (req, res) => res.status(200).send());
app.post("/notify", (req, res) => {
    notify(req.body, reply => res.send(reply));
});

app.listen(port, () => console.log(`server is running on port: ${port}`));

const notify = ({title, message}, callBack) => {
    notifier.notify(
        {
            title: title || "No title",
            message: message || "No message",
            icon: path.join(__dirname, "gophertuts.png"),
            sound: true,
            wait: true,
            reply: true,
            closeLabel: "Completed ?",
            timeout: 15
        },
        (err, response, reply) => {
            callBack(reply);
        }
    );
}