// Chayapol Wongpuwarak 65070501011 
public abstract class Teacher {

	// Delete void Teach

	private void makeAnnouncements() {
		System.out.println("made announcements..");
	}

	private void takeAttendence() {
		System.out.println("took attendence..");
	}

	private void collectPaperWork() {
		System.out.println("collected paperwork..");
	}

	private void conductHallwayDuties() {
		System.out.println("conducted hallway duties..");
	}

	public void performOtherResponsibilities() {
        makeAnnouncements();
        takeAttendence();
        collectPaperWork();
        conductHallwayDuties();
    }
}
