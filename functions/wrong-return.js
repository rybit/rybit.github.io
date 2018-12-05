
exports.handler = function(event, context, callback) {
    console.log({"event": event, "context": context})

    callback(null, {
    statusCode: 200,
    body: 1234541
    });
}
