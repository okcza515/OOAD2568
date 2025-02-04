
public class ReportGenerator {
	
	private ReportTransaction reportObject;
	
	public void generateReport(){
		System.out.println(reportObject.getName()
		+" "
		+reportObject.productBreakDown()
		+" "
		+reportObject.getDate());
	}

}
