class WorkoutPage {
  constructor() {}

  public showTab(tabName: string) {
    const tabs = document.getElementById("intervals-card");
    // Hide all tabs
    tabs.querySelectorAll(".tab-content").forEach((tab) => {
      tab.classList.add("hidden");
    });
    tabs.querySelectorAll(".tab-button").forEach((button) => {
      button.classList.remove("active");
    });

    // Show selected tab
    document.getElementById(tabName + "-tab").classList.remove("hidden");
    if (event.target instanceof HTMLElement) {
      const button =
        event.target.nodeName === "BUTTON"
          ? event.target
          : event.target.closest("button");
      button?.classList.add("active");
    }
  }

  public selectInterval(from: number, to: number) {}

  public selectSegment(name: string, from: number, to: number) {
    this.updateSelectionIndicator(name);
  }

  public selectClimb(name: string, from: number, to: number) {
    this.updateSelectionIndicator(name);
  }

  private updateSelectionIndicator(text: string) {
    document.getElementById("selection-indicator").textContent = text;
  }
}

globalThis.workoutPage = new WorkoutPage();

// Initialize with default selection
window.addEventListener("load", function () {
  globalThis.workoutPage.updateSelectionIndicator("Interval 2 (1.0km)");
});
