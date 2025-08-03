import { Calendar } from "fullcalendar";

class ProfileCalendar extends HTMLElement {
  constructor() {
    super();

    this.apiWorkoutsCalendarRoute = this.getAttribute(
      "api-workouts-calendar-route",
    );
  }

  connectedCallback() {
    var calendar = new Calendar(this, {
      timeZone: this.getAttribute("timezone"),
      initialView: "dayGridMonth",
      locale: Intl.DateTimeFormat().resolvedOptions().locale,
      firstDay: 1,
      aspectRatio: 2,
      eventContent: function (arg) {
        let eventSpan = document.createElement("div");

        eventSpan.innerHTML = arg.event.title;
        eventSpan.classList.add("px-2", "overflow-hidden");

        return { domNodes: [eventSpan] };
      },
      events: {
        url: this.apiWorkoutsCalendarRoute,
        display: "display",
        failure: function () {
          console.log("there was an error while fetching events!");
        },
        success: function (response) {
          return response.results;
        },
      },
    });
    calendar.render();
  }
}

customElements.define("profile-calendar", ProfileCalendar);
