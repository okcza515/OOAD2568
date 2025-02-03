public class Teacher {

    private void makeAnnoucements() {
        System.out.println("The teacher is making an annoucement");
    }

    private void takeAttendance() {
        System.out.println("The teacher is taking an attendance");
    }

    private void collectPaperWork() {
        System.out.println("The teacher is collecting paperwork");
    }

    private void conductHallWayDuties() {
        System.out.println("The teacher is conducting hallway duties");
    }

    public void performOtherResponsibilities() {
        makeAnnoucements();
		takeAttendance();
		collectPaperWork();
		conductHallWayDuties();
        System.out.println("The teacher is performing other responsibilities");
    }

}

//Supanut Wongtanom 65070503437
