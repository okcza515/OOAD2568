
public abstract class Teacher {
	// delete void teach
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

	public void performOtherResponsibilities() {// why recursive this brooooooo
		makeAnnouncements();
		takeAttendence();
		collectPaperWork();
		conductHallwayDuties();
		performOtherResponsibilities();
	}
}

// 65070501065 Kamolpop Poonsawat
