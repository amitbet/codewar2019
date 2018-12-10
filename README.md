# Codewar 2018
![code war image](https://github.com/codearmada/codewar2018/blob/master/sample.gif "code war!")

A CodeWar bot game based on Conway's game of life.

## Installing the developer environment:
* [Download NodeJS](https://nodejs.org/en/download/)
* Install [git](https://git-scm.com/downloads)
* In cmdline/bash run:

  **git clone https://github.com/amitbet/codewar2018.git**
* Install node-static globally (install.cmd or install.sh script)
* Run the (run.cmd / run.sh) script
* Open (http://localhost:60606)
* Open any Js code editor and modify code/my-bot.js until its **fit for war!**

This code war game utilizes [conway's Game Of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) which is an celular automata
where the next position of the board is calculated by subjecting the current board to 4 simple [rules](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life#Rules)

The objective of this game is to send your pixels to the other side of the screen, scoring a hit and decreasing your opponent's life force, while defending your side against such invaders.

## Tips
* Use Golly to create patterns and copy them into code (instructions below)
* If you wish for simpler code, the bots contain a more mundane implementation that puts pixels without using RLE format.
* Do not just use my-bot.js code as is, **you can improve every aspect of it**
* Search for helpful patters on [lifewiki site](http://www.conwaylife.com/wiki/Category:Patterns)

## Golly:
* Download and install [golly](https://sourceforge.net/projects/golly/files/golly/golly-3.2/)

  Golly is a studio where you can experiment with game of life patterns and see how they develop and interact
* I have included a nice RLE to pixels routine to make it easier to import patterns from Golly/LifeWiki, you can just copy the shape in golly and paste as string into your code.

![Golly Image](https://github.com/amitbet/codewar2018/blob/master/golly.jpg "golly")
**Now just paste it into your code to get:**
x = 4, y = 4, rule = B3/S23
o$bo$2bo$3bo!

Use the last bit as a string that will represent the pattern:
var diagonalLine = "o$bo$2bo$3bo!"; // diagnoal line cost=4 pixels
pixels = tryPlaceRle(data, diagonalLine, 4);


