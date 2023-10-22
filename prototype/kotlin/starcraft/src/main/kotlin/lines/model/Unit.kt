package lines.model

interface Unit {
    fun GetMineralCapacity(): Int
    fun SetMineralCapacity(accumulateMineral: Int)

    fun Harvest()
    fun Attack()
    fun Building()
}