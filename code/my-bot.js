//INPROVE THIS CODE

(function () {

    var spaceShip = "3o$o2bo$o$o$bobo!"; //lightweight spaceShip 8
    var lGlider = "3o$o$bo!"; //glider 5
    var rGlider = "3o$2bo$bo!"; //glilder 5
    var preblock = "2o$o!"; // pre block 3
    var zap = "o$2o$bo!"; // creates a boat(6) costs:4
    var blinker = "3o!";


    function getRnd(min, max) {
        return Math.floor(Math.random() * (max - min + 1)) + min;
    }

    function registerArmy() {
        window.registerArmy({
            name: 'Power-Hog',
            icon: 'gladiator',
            cb: cb
        });
    }

    function getRandom() {
        return Math.random();
    }
    setTimeout(registerArmy, 0);

    var plan = {
        ongoing: [
            'spaceShip',
        ]
    };


    function cb(data) {
        var pixels = [];
        var plan;
        if (data.generation === 1) {
            planIndex = 0;
            fenceLocation = 0;
            fenceRow = 15;
        }
        if (data.generation < 200) {
            plan = ['fence', 'spaceship'];
        } else if (data.generation < 440) {
            plan = ['fence'];
        } else {
            plan = ['spaceship'];
        }
        planIndex = planIndex % plan.length;
        if (plan[planIndex] === 'mine') {
            pixels = tryPlaceRle(data, preblock, 3);
        } else if (plan[planIndex] === 'fence') {
            pixels = tryPlaceFence(data);
        } else if (plan[planIndex] === 'glider') {
            pixels = tryPlaceRle(data, rGlider, 5);
        } else if (plan[planIndex] === 'spaceship') {
            pixels = tryPlaceRle(data, spaceShip, 9);
        }

        if (pixels.length > 0) {
            planIndex = (planIndex + 1) % plan.length;
        }
        return pixels;
    };

    function tryPlaceFence(data, col, row) {
        var pixels = [];
        var r, c;
        c = col || fenceLocation;
        r = row || data.rows - fenceRow;
        pixels = tryPlaceRle(data, blinker, 3, c, r);

        if (pixels.length > 0)
            fenceLocation += 10;

        if (fenceLocation > data.cols - 2) {
            fenceLocation = 0;
            fenceRow += 10;
        }
        return pixels;
    }

    function tryPlaceRle(data, rle, neededBudget, col, row) {
        var pixels = [];

        if (data.budget >= neededBudget) {
            c = (col === 0) ? 0 : col || getRnd(0, data.cols - 2);
            r = (row === 0) ? 0 : row || getRnd(20, 80);
            pixels = getPixelsFromRle(rle, c, r);
        }
      
        return pixels;
    }


    function getPixelsFromRle(rle, c, r, pixels) {
        var pixels = [];
        var num = '';
        var x = 0;
        var y = 0;
        var l;

        for (var s in rle) {
            var s = rle[s];
            if (s === 'b') {
                x = num === '' ? x + 1 : x + parseInt(num);
                num = '';
            } else if (s === 'o') {
                var i = num === '' ? 1 : parseInt(num);
                while (i--)
                    pixels.push([c + x + i, r + y]);

                x = num === '' ? x + 1 : x + parseInt(num);
                num = '';
            } else if (s === '$') {
                y += num === '' ? 1 : parseInt(num);
                x = 0;
                num = '';
            } else if (s === '!')
                break;
            else if (parseInt(s).toString() !== 'NaN') {
                num += s;
            }
        }
        return pixels;
    };

})();
