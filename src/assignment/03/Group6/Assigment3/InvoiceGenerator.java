// ✅ คลาสสำหรับสร้างใบแจ้งหนี้ (Invoice)
// 65070501025 ธเนศ จอมพูล
class InvoiceGenerator {

    public void prepareInvoice(TransactionDetails transaction) {
        System.out.println("🔹 Generating Invoice for: " + transaction.getName());
        System.out.println("Date: " + transaction.getDate());
        System.out.println("Product Breakdown: " + transaction.getProductBreakdown());
        System.out.println("✅ Invoice Generated!");
    }
}
