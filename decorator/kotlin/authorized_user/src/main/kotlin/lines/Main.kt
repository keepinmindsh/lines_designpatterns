package lines

fun main() {
    val switAdmin = SwitAdmin(SwitUser())

    val authorizedUser = switAdmin.GetAuthorizedUser()

    print(authorizedUser)
}

data class UserModel(
    val UserId: String,
    val UserName: String,
    var Tel: String,
    var Address: String,
    var Role: String
)

interface User {
    fun GetUser(): List<UserModel>
}

interface AuthrizedUser {
    fun GetAuthorizedUser(): List<UserModel>
}


class SwitUser: User {
    override fun GetUser(): List<UserModel> {
        return ArrayList<UserModel>()
    }
}

class SwitAdmin(val user: User) : AuthrizedUser {
    override fun GetAuthorizedUser(): List<UserModel> {

        val userList = user.GetUser()

        val filterUserList = userList.filter { item -> checkRole(item.Role) }

        return filterUserList
    }

    private fun checkRole(role:String): Boolean {
        return role == "Admin"
    }
}
