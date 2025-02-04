public class Main {
    public static void main(String[] args) {
        // สร้างข้อมูลธุรกรรม
        TransactionDetails transaction = new TransactionDetails("John Doe", "2025-02-03", "Laptop, Mouse"); //test
        System.out.println("Transaction Details: " + transaction.getName() + ", " + transaction.getDate() + ", " + transaction.getProductBreakdown());
        // เรียกใช้งาน CustomerTransaction

        // CustomerTransaction customerTransaction = new CustomerTransaction(transaction);
        // customerTransaction.processTransaction();
    }
}