// ✅ คลาสหลักสำหรับเก็บข้อมูลธุรกรรม
// 65070501025 ธเนศ จอมพูล
class TransactionDetails {
    private String name;
    private String date;
    private String productBreakdown;

    public TransactionDetails(String name, String date, String productBreakdown) {
        this.name = name;
        this.date = date;
        this.productBreakdown = productBreakdown;
    }

    public String getName() {
        return name;
    }

    public String getDate() {
        return date;
    }

    public String getProductBreakdown() {
        return productBreakdown;
    }
}