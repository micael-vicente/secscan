package com.somepackage;

//imports

public class SomeRepository {
  public List<AccountDTO> findAccountsByCustomerId(String customerId) throws SQLException {

    String sql = "select * from Accounts where customer_id = '%s'";

    String sql = "select "
        + "customer_id, acc_number,branch_id,balance "
        + "from Accounts where customer_id = '%s'";

    Connection c = dataSource.getConnection();
    ResultSet rs = c.createStatement().executeQuery(String.format(sql, customerId));

    //...

  }

  public AccountDTO findAccountById(String id) throws SQLException {

    String sql = "select "
        + "customer_id, acc_number,branch_id, balance "
        + "from Accounts where id = '%s'";

    Connection c = dataSource.getConnection();
    ResultSet rs = c.createStatement().executeQuery(String.format(sql, id));

    //...

  }
}