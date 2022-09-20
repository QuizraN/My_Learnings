"use strict";
// ⏇⏇⏇⏇⏇⏇⏇⏇⏇⏇⏇⏇⏇⏇⏇⏇⏇⏇⏇⏇⏇⏇⏇⏇⏇⏇⏇⏇
//   Exercise 4 – Types
// ⏈⏈⏈⏈⏈⏈⏈⏈⏈⏈⏈⏈⏈⏈⏈⏈⏈⏈⏈⏈⏈⏈⏈⏈⏈⏈⏈⏈
exports.__esModule = true;
// Objectives:
// • Understand how TypeScript performs code flow analysis
// • Create and apply union and intersection types
// • Use basic type guards (narrowing types w/ typeof, instanceof, etc.)
exports["default"] = (function () {
    // ======== Exercise 4.0 ========
    // TypeScript is intelligent about the possible types of a variable,
    // depending on the code path.
    // Instructions:
    // • Simply inspect the possible types by hovering over `text` to see
    //   how the inferred type changes if assumptions can be safely made
    //   about the possible types within the given code path.
    function trimmedLength1(text) {
        text; // text: string | null | undefined
        if (typeof text === 'string') {
            text; // text: string
            return text.trim().length;
        }
        text; // text: null | undefined
    }
    function trimmedLength2(text) {
        text; // text: string | null | undefined
        if (typeof text === 'string') {
            text; // text: string
            return text.trim().length;
        }
        else if (text == null) {
            text; // text: null | undefined (remember == coerces undefined)
            return;
        }
        text; // text: never
    }
    function trimmedLength3(text) {
        if (text) {
            return text.trim().length;
        }
        text; // text: string | null | undefined (because '' == false)
    }
    function trimmedLength4(text) {
        if (!text) {
            text;
            return;
        }
        return text.trim().length; // text: string
    }
    function trimmedLength5(text) {
        text; // text: any
        if (typeof text === 'string') {
            return text.trim().length; // text: string
        }
        text; // text: any (note how TS does not subtract types from `any`)
    }
    console.log('[Exercise 4.0]', "".concat(trimmedLength1("   hi     ")));
    // ======== Exercise 4.1 ========
    // Instructions:
    // • Restrict type of `value` to `string OR number`
    // • Fix any resulting errors.
    console.log('[Exercise 4.1]');
    function doStuff(value) {
        if (typeof value === 'string') {
            console.log(value.toUpperCase().split('').join(' '));
        }
        else if (typeof value === 'number') {
            console.log(value.toPrecision(5));
        }
        value; // hover over `value` here
    }
    doStuff(2);
    doStuff(22);
    doStuff(222);
    doStuff('hello');
    doStuff('true');
    doStuff('{}');
    // ======== Exercise 4.2 ========
    // Instructions:
    // • Use a type guard to fill out the body of the `padLeft` function.
    function padLeft(value, padding) {
        if (typeof padding === 'string') {
            return padding + value;
        }
        else {
            return "".concat(Array(padding + 1).join(' ')).concat(value);
        }
        // if padding is a number, return `${Array(padding + 1).join(' ')}${value}`
        // if padding is a string, return padding + value
    }
    console.log('[Exercise 4.2]', "\n    ".concat(padLeft('🐕', 0), "\n    ").concat(padLeft('🐕', '🐩'), "\n    ").concat(padLeft('🐕', '🐩🐩'), "\n    ").concat(padLeft('🐕', '🐩🐩🐩'), "\n    ").concat(padLeft('🐕', '🐩🐩🐩🐩'), "\n  "));
    // ======== Exercise 4.3 ========
    // Instructions:
    // • Add type annotations (`any` excluded)
    // • Inspect inferred type of `element` in different code branches
    // • Bonus: turn `flatten` into a generic function
    var numbers = [1, 2, 3, [44, 55], 6, [77, 88], 9, 10];
    function flatten(array) {
        var flattened = [];
        for (var _i = 0, array_1 = array; _i < array_1.length; _i++) {
            var element = array_1[_i];
            if (Array.isArray(element)) {
                element; // any[]
                flattened.push.apply(// any[]
                flattened, element);
            }
            else {
                element; // any
                flattened.push(element);
            }
        }
        return flattened;
    }
    var flattenedNumbers = flatten(numbers);
    console.log('[Exercise 4.3]', flattenedNumbers);
    var Bird = /** @class */ (function () {
        function Bird(species) {
            this.species = species;
        }
        Bird.prototype.layEggs = function () {
            console.log('Laying bird eggs.');
        };
        Bird.prototype.fly = function (height) {
            console.log("Flying to a height of ".concat(height, " meters."));
        };
        return Bird;
    }());
    ;
    var Fish = /** @class */ (function () {
        function Fish(species) {
            this.species = species;
        }
        Fish.prototype.layEggs = function () {
            console.log('Laying fish eggs.');
        };
        Fish.prototype.swim = function (depth) {
            console.log("Swimming to depth of ".concat(depth, " meters."));
        };
        return Fish;
    }());
    function getRandomAnimal() {
        var animals = [
            new Bird('puffin'),
            new Bird('kittiwake'),
            new Fish('sea robin'),
            new Fish('leafy seadragon'),
        ];
        return animals[Math.floor(Math.random() * animals.length)];
    }
    function interrogateAnimal(animal) {
        if (animal === void 0) { animal = getRandomAnimal(); }
        console.log("instanceof", animal instanceof Bird, animal.species);
        if (animal instanceof Fish) {
            (animal).swim(10);
        }
        else {
            (animal).fly(10); // call only if it is a bird
        }
        // animal.swim(10) // call only if it is a fish
        return animal.species;
    }
    console.log('[Exercise 4.4]', "We've got a ".concat(interrogateAnimal(), " on our hands!"));
})();
