<script>
  import Codehighlighter from "$lib/components/page/Codehighlighter.svelte";
  import Sandbox from "$lib/components/page/Sandbox.svelte";
  import highlightJs from "highlight.js/lib/languages/javascript";
  import highlightCss from "highlight.js/lib/languages/css";
  import highlightXml from "highlight.js/lib/languages/xml";
  let script = `(() => {
  const xobox = [];
  const x = 'x';
  const o = 'o';
  let turn = x;
  const steps = [];
  let game_over = false;
  const grid = document.getElementById('game-board');
  const message = document.getElementById('message');
  const restart = document.getElementById('restart');

  for (let i = 0; i < 3; i++) {
    xobox.push([]);
    for (let j = 0; j < 3; j++) {
      xobox[i].push(document.getElementById(i.toString() + j.toString()));
    }
  }

  const check = () => {
    if (xobox[0][0].innerText != '' && xobox[0][0].innerText == xobox[0][1].innerText && xobox[0][0].innerText == xobox[0][2].innerText) {
      return true;
    }
    if (xobox[1][0].innerText != '' && xobox[1][0].innerText == xobox[1][1].innerText && xobox[1][0].innerText == xobox[1][2].innerText) {
      return true;
    }
    if (xobox[2][0].innerText != '' && xobox[2][0].innerText == xobox[2][1].innerText && xobox[2][0].innerText == xobox[2][2].innerText) {
      return true;
    }
    if (xobox[0][0].innerText != '' && xobox[0][0].innerText == xobox[1][0].innerText && xobox[0][0].innerText == xobox[2][0].innerText) {
      return true;
    }
    if (xobox[0][1].innerText != '' && xobox[0][1].innerText == xobox[1][1].innerText && xobox[0][1].innerText == xobox[2][1].innerText) {
      return true;
    }
    if (xobox[0][2].innerText != '' && xobox[0][2].innerText == xobox[1][2].innerText && xobox[0][2].innerText == xobox[2][2].innerText) {
      return true;
    }
    if (xobox[0][0].innerText != '' && xobox[0][0].innerText == xobox[1][1].innerText && xobox[0][0].innerText == xobox[2][2].innerText) {
      return true;
    }
    if (xobox[2][0].innerText != '' && xobox[2][0].innerText == xobox[1][1].innerText && xobox[2][0].innerText == xobox[0][2].innerText) {
      return true;
    }
    return false;
  };

  const getter = (i, j) => xobox[i][j].innerText;

  const setter = (i, j, value) => {
    xobox[i][j].innerText = value;
  };

  const reset = () => {
    restart.blur();
    grid.focus();
    steps.length = 0;
    game_over = false;
    message.innerText = 'Best of Luck!';
    for (let i = 0; i < 3; i++) {
      for (let j = 0; j < 3; j++) {
        setter(i, j, '');
      }
    }
    turn = x;
  };

  const predict_winner = () => {
    if (check()) {
      message.innerText = "Congratulations! '" + turn + "' won.";
      game_over = true;
      return true;
    }
    if (steps.length == 9) {
      message.innerText = "It's a tie!";
      game_over = true;
      return true;
    }
    game_over = false;
    return false;
  };

  const next = (i, j) => {
    if (getter(i, j) == '') {
      if (predict_winner()) {
        return;
      }
      setter(i, j, turn);
      steps.push([i, j, turn]);
      if (predict_winner()) {
        return;
      }
      if (turn == x) {
        turn = o;
      } else {
        turn = x;
      }
    }
  };

  const prev = () => {
    if (!game_over) {
      if (steps.length > 0) {
        const temp = steps.pop();
        setter(temp[0], temp[1], '');
        turn = temp[2];
      }
    }
  };

  restart.onclick = () => {
    reset();
  };

  for (let i = 0; i < 3; i++) {
    for (let j = 0; j < 3; j++) {
      (() => {
        xobox[i][j].onclick = () => {
          grid.focus();
          next(i, j);
        };
        xobox[i][j].ondblclick = () => {
          grid.focus();
          prev();
        };
      })();
    }
  }
})();
`;
  let css = `:root {
  --color-1: #149414;
  --color-2: #9ccc9c;
  --color-3: #2b5329;
  --color-4: #ff0000;
}

.game-board {
  margin: 1em;
  width: max-content;
  display: grid;
  grid-template-rows: 50px 50px 50px;
  grid-template-columns: 50px 50px 50px;
  grid-gap: 10px;
}

.box {
  display: flex;
  justify-content: center;
  align-items: center;
  border: 5px solid var(--color-2);
  color: var(--color-4);
  font-size: 20px;
  background-color: var(--color-3);
}

.box:hover {
  background-color: var(--color-1);
}

button {
  margin: 1em;
  border-radius: 5px;
  background-color: var(--color-3);
  color: white;
  padding: 5px;
  border: 2px solid white;
}

button:hover {
  background-color: white;
  color: var(--color-3);
  border: 2px solid var(--color-3);
}

.center {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  text-align: center;
}

#message {
  margin: 1em;
  width: max-content;
  border-radius: 5px;
  padding: 5px;
  background-color: var(--color-2);
  border: 2px solid var(--color-3);
  color: var(--color-3);
}
`;
  let html = `&lt;html&gt;

&lt;head&gt;
  &lt;link rel='stylesheet' href='style.css'&gt;
  &lt;title&gt;
    TicTacToe
  &lt;/title&gt;
&lt;/head&gt;

&lt;body&gt;
  &lt;div class=&quot;center&quot;&gt;
    &lt;div class=&quot;game-board&quot; id='game-board'&gt;
      &lt;div class=&quot;box&quot; id='00'&gt;&lt;/div&gt;
      &lt;div class=&quot;box&quot; id='01'&gt;&lt;/div&gt;
      &lt;div class=&quot;box&quot; id='02'&gt;&lt;/div&gt;
      &lt;div class=&quot;box&quot; id='10'&gt;&lt;/div&gt;
      &lt;div class=&quot;box&quot; id='11'&gt;&lt;/div&gt;
      &lt;div class=&quot;box&quot; id='12'&gt;&lt;/div&gt;
      &lt;div class=&quot;box&quot; id='20'&gt;&lt;/div&gt;
      &lt;div class=&quot;box&quot; id='21'&gt;&lt;/div&gt;
      &lt;div class=&quot;box&quot; id='22'&gt;&lt;/div&gt;
    &lt;/div&gt;
    &lt;div id='message'&gt;Best of Luck!&lt;/div&gt;
    &lt;button id='restart'&gt;Restart&lt;/button&gt;
  &lt;/div&gt;

  &lt;script src=&quot;index.js&quot;&gt;&lt;/script&gt;
&lt;/body&gt;

&lt;/html&gt;
`;
  html = html
    .replaceAll("&lt;", "<")
    .replaceAll("&gt;", ">")
    .replaceAll("&quot;", '"')
    .replaceAll("&#39;", "'")
    .replaceAll("&amp;", "&");
</script>

<h1>TicTacToe</h1>

<h2>Source Code</h2>
<Codehighlighter
  code={script}
  langDef={highlightJs}
  langName={"javascript"}
  fileName={"./index.js"}
/>
<Codehighlighter
  code={css}
  langDef={highlightCss}
  langName={"css"}
  fileName={"./style.css"}
/>
<Codehighlighter
  code={html}
  langDef={highlightXml}
  langName={"html"}
  fileName={"./index.html"}
/>

<h2>Output</h2>
<Sandbox title="tic-tac-toe" src="/resources/blog/tictactoe/index.html" />
