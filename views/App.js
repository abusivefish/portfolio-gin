import React from 'react';
import './App.css';

//global header

function App() {
  return (
    <div className="App">
      <header className="navbar">
        <a href="https://github.com/abusivefish">
        <img 
            className="GH" 
            alt="" 
            src="./GitHub-Mark-32px.png" 
            id="gh-ref"
        />
        </a>
          <div className="dropdown">
            My Projects
          </div>
          <img

          />
        
      </header>
    
      <div className="carousel">
        <div className="btn">
        </div>
      </div>

      <footer>
        
      </footer>
    </div>
  );
}


/*
function InvertColor() {
  do thing
}


Old Brick 	#912133 	Primary
Baltic Sea 	#232127 	Info
Dingley 	#618447 	Success
Christine 	#de740f 	Warning
Pomegranate 	#f44336 	Danger


*/


export default App;
