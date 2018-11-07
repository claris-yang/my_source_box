package hello;

public interface UserMapper {
    public int insertUser(UserBean user) throws Exception;

    public int updateUser(UserBean user, int id) throws Exception;

    public int deleteUser(int id) throws Exception;

    public UserBean selectUserById(int id) throws Exception;

}
