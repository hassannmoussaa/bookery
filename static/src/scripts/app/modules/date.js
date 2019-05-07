require(['jquery', 'moment'], function ($, moment) {
    var date = $('#date');
    if (date.length > 0) {
        function update() {
            date.html(moment().utcOffset(3).format('ddd D MMMM YYYY hh:mm:ss a'));
        }
        setInterval(update, 1000);
    }
});