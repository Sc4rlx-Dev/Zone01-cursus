const load = document.getElementById("load")
const dis = document.getElementById("rm")
const imag = document.getElementById("image")


load.addEventListener('click' , async() => {
    const url = 'https://picsum.photos/600/400'
    const r = await fetch(url)

    const img = document.createElement('img')
    img.src = r.url
    img.style.maxWidth = '1000px'

    imag.innerHTML = ''
    imag.appendChild(img)
})




'use strict';

process.stdin.resume();
process.stdin.setEncoding('utf-8');

let inputString = '';
let currentLine = 0;

process.stdin.on('data', function(inputStdin) {
    inputString += inputStdin;
});

process.stdin.on('end', function() {
    inputString = inputString.split('\n');

    main();
});

function readLine() {
    return inputString[currentLine++];
}


