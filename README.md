# codewar2018
code war based on game of life
to install the developer environment: 
* [Download NodeJS](https://nodejs.org/en/download/)
* install node-static globally (install.cmd or install.sh script)
* run the (run.cmd / run.sh) script
* open (http://localhost:60606)
* open any Js code editor and modify code/my-bot.js until its **fit for war!**

This code war game utilizes [conway's Game Of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) which is an celular automata
where the next position of the board is calculated by subjecting the current board to 4 simple [rules](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life#Rules)

The objective of this game is to send your pixels to the other side of the screen, scoring a hit and decreasing your opponent's life force, While defending your side against such invaders.

**This is the game:**

![code war image](https://github.com/codearmada/codewar2018/blob/master/sample.gif "code war!")

**Be sure to use these tips:**
* download and install [golly](https://sourceforge.net/projects/golly/files/golly/golly-3.2/)
* search for patters on [lifewiki site](http://www.conwaylife.com/wiki/Category:Patterns)
* I have included a nice RLE to pixels routine to make it easier to import patterns from Golly/LifeWiki, you can just copy the shape in golly and paste as string into your code.
