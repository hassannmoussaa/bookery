var regions = ["beirut", "mount lebanon", "beqaa", "nabatieh", "north", "south"];

for (var i = 0; i < regions.length; i++) {
    var data = { name: regions[i] };
    var request = new Request().init(_APIURL + '/regions', data);
    request.onSuccess(function (result) {
        console.log(result);
    });
    request.onFailure(function (err) {
        console.error(err.data.message);
    });
    request.post();
}