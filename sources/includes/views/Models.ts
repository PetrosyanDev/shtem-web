class Question {
    public shtemaran: string
    public bajin: number
    public mas: number
    public number: number
    public text: string
    public options: string[]
    public answers: number[]

    constructor(shtemaran: string, bajin: number, mas: number, q_number: number, text: string, options: string[], answers: number[]) {
        this.shtemaran = shtemaran
        this.bajin = bajin
        this.mas = mas
        this.number = q_number
        this.text = text
        this.options = options
        this.answers = answers
    }
}

export default Question