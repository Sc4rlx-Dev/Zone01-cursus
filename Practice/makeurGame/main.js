
const bt1 = document.getElementById("start")
const bt2 = document.getElementById("end")
const imag = document.getElementById("image")
const Game_started = false



const src = document.getElementById('header')

bt1.addEventListener('click', () => {
    if (!document.querySelector('#header img')) {
        alert('Game Started')
        const img = document.createElement('img')
        img.src = "./img.png"
        img.style.width = '350px'
        src.innerHTML = ''
        src.appendChild(img)
        Game_started = true
    }
})

function close() {
    bt2.addEventListener('click', () => {
    src.innerHTML = ''
})
}

if(Game_started == true){
    const newGame = document.getElementById("new")
    const startgame = document.getElementById("display")



    newGame.addEventListener('click' , ()=> {
        const window = document.createElement('img2')
        img2.src = "./phone.jpg"
        const phone = getElementById('header')
        phone.removeChild(src)
    })
}


                                                                          