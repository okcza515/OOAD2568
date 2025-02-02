public class Teacher{
    private String name;
    
    private void makeAnnouncements(){
        System.out.println("This announcement has been announced <3");
    }

    private void takeAttendance(){
        System.out.println("Attendance has been attended");
    }

    private void collectPaperWork(){
        System.out.println("PaperWork has been papered");
    }

    private void conductHallwayDuties(){
        System.out.println("Hallway duties has been dutied");
    }

    public void performOtherResponsibilities(){
        makeAnnouncements();
        takeAttendance();
        collectPaperWork();
        conductHallwayDuties();
        // System.out.println("OtherResponsibilities has been Responded");
    }
}
//Pontkorn Wichaporn 65070503427