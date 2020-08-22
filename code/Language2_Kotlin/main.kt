import java.util.Random
 
val rand = Random(0)
 
class Field(val w: Int, val h: Int) 
{
    val s = List(h) { BooleanArray(w) }
    operator fun set(x: Int, y: Int, b: Boolean) 
    {
        s[y][x] = b
    }
 
    fun next(x: Int, y: Int): Boolean 
    {
        var on = 0
        for (i in -1..1) {
            for (j in -1..1) 
            {
                if (state(x + i, y + j) && !(j == 0 && i == 0)) on++
            }
        }
        return on == 3 || (on == 2 && state(x, y))
    }
 
    fun state(x: Int, y: Int): Boolean 
    {
        if ((x !in 0 until w) || (y !in 0 until h)) return false
        return s[y][x]
    }
}
 
class gameOfLife() 
{
    val w: Int
    val h: Int
    var a: Field
    var b: Field
 
    init 
    {
      w = 5
      h = 5
      a = Field(w, h)
      b = Field(w, h)
      for (i in 0 until w * h) 
      {
          a[rand.nextInt(w), rand.nextInt(h)] = true
      }
    }

    fun step() 
    {
        for (y in 0 until h) 
        {
            for (x in 0 until w) 
            {
                b[x, y] = a.next(x, y)
            }
        }
        val t = a
        a = b
        b = t
    }
 
    override fun toString(): String 
    {
        val sb = StringBuilder()
        for (y in 0 until h) 
        {
            for (x in 0 until w) 
            {
                val c = if (a.state(x, y)) '#' else '.'
                sb.append(c)
            }
            sb.append('\n')
        }
        return sb.toString()
    }
}
 
fun main(args: Array<String>) 
{
    val lives = listOf(gameOfLife())
    val generations = 4
    for(game in lives)
    {
      repeat(generations + 1)
      {
        println("Generation: $it\n$game")
        game.step()
      }
    }
}