import { lerp } from './utils.js';
class Road {
    constructor( x, width, laneCount = 3) {
        this.width = width
        this.laneCount = laneCount
        this.x = x
        
        this.left = x -this.width/2
        this.right = x + this.width/2

        const infinity = 1000000
        this.top = -infinity
        this.bottom = infinity 
        
        const topleft = {x: this.left, y: this.top}
        const topright = {x: this.right, y: this.top}
        const bottomleft = {x: this.left, y: this.bottom}
        const bottomright = {x: this.right, y: this.bottom}
        
        this.borders = [
            [topleft,bottomleft],
            [topright,bottomright]  
        ]
    }

    getLaneCenter(laneIndex) {
        const laneWidth = this.width / this.laneCount
        return (
            this.left + 
            laneWidth / 2 + 
            Math.min(laneIndex, this.laneCount - 1) * laneWidth
        )
    }

    draw(ctx) {
        ctx.lineWidth = 5
        ctx.strokeStyle = 'white'

        for(let i=1; i <= this.laneCount-1; i++) {
            const x = lerp(
                this.left, 
                this.right, 
                i / this.laneCount
            )
           
            ctx.setLineDash([20, 20])           
            ctx.beginPath()
            ctx.moveTo(x, this.top)
            ctx.lineTo(x, this.bottom)
            ctx.stroke()
        } 
        
        ctx.setLineDash([])
        this.borders.forEach(border => {
            ctx.beginPath()
            ctx.moveTo(border[0].x, border[0].y)
            // ctx.lineTo(border[].x + border.width, border.y)
            // ctx.lineTo(border.x + border.width, border.y + border.height)
            // ctx.lineTo(border.x, border.y + border.height)
            ctx.lineTo(border[1].x, border[1].y)
            ctx.stroke()
        })
    }

    
}


export { Road }