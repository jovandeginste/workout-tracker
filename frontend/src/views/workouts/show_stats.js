import ApexCharts from 'apexcharts';

class WorkoutStats extends HTMLElement {
  constructor() {
    super();
    this.style.display = 'block';
  }

  connectedCallback() {
    this.mapElement = document.getElementById(this.getAttribute('map-id'));
    this.preferredUnits = JSON.parse(this.getAttribute('preferred-units'));
    this.data = JSON.parse(this.getAttribute('data'));
    this.tz = this.getAttribute('tz');
    this.lang = this.getAttribute('lang');
    this.translations = JSON.parse(this.getAttribute('translations'));

    let theme = 'light';
    if (window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches) {
      theme = 'dark';
    }

    let options = {
      theme: { mode: theme },
      chart: {
        height: 400,
        animations: { enabled: false },
        toolbar: { show: false },
      },
      legend: {
        position: 'top',
        formatter: (seriesName, opts)=>{
          if(opts.seriesIndex>3) return '';
          return seriesName;
        },
        markers: { size: [12,12,12,12,0] }
      },
      tooltip: {
        x: { format: 'HH:mm', },
        y: [
          {
            formatter: (val, opts) => {
            let p = this.data[opts.dataPointIndex]
            let el = document.createElement('div');
            el.setAttribute("data-lat", p.Item.firstPoint.lat);
            el.setAttribute("data-lng", p.Item.firstPoint.lng);
            el.setAttribute("data-title", p.Label);

            this.mapElement.setMarker(el)
            return val + " " + this.preferredUnits.speed;
            }
          },
          { formatter: (val, opts) => { return val + " " + this.preferredUnits.elevation; } },
          { formatter: (val, opts) => { return val + " " + this.preferredUnits.heartRate; } },
          { formatter: (val, opts) => { return val + " " + this.preferredUnits.cadence; } },
          { formatter: (val, opts) => { return val + " " + this.preferredUnits.distance; } },
          { formatter: (val, opts) => { return formatDuration(val); } },
        ],
      },
      stroke: {
        width: 2,
        curve: 'smooth',
      },
      markers: {
        size: 1,
      },
      series: [
        {
          name: this.translations.averagespeed,
          type: "line",
          data: this.data.map(e => ({ x: e.Item.firstPoint.time, y: e.Item.localAverageSpeed })),
        },
        {
          name: this.translations.elevation,
          type: "area",
          data: this.data.map(e => ({ x: e.Item.firstPoint.time, y: e.Item.localElevation })),
        },
        {
          name: this.translations.heartrate,
          type: "line",
          display: false,
          data: this.data.map(e => ({ x: e.Item.firstPoint.time, y: e.Item.localHeartRate })),
        },
        {
          name: this.translations.cadence,
          type: "line",
          display: false,
          data: this.data.map(e => ({ x: e.Item.firstPoint.time, y: e.Item.localCadence })),
        },
        {
          name: this.translations.distance,
          type: "none",
          data: this.data.map(e => ({ x: e.Item.firstPoint.time, y: e.Item.localTotalDistance })),
        },
        {
          name: this.translations.duration,
          type: "none",
          data: this.data.map(e => ({ x: e.Item.firstPoint.time, y: e.Item.totalDurationSeconds })),
        },
      ],
      xaxis: {
        labels: {
          formatter: (val, ts, opts) => {
            return new Date(ts).toLocaleTimeString(this.lang, { timeZone: this.tz })
          },
        },
        type: "datetime",
      },
      yaxis: [
        {
          min: 0,
          labels: {
            formatter: (val) => {
              return val + " " + this.preferredUnits.speed;
            },
          },
        },
        {
          labels: {
            formatter: (val) => {
              return val + " " + this.preferredUnits.elevation;
            },
          },
          opposite: true,
        },
        {
          labels: {
            formatter: (val) => {
              return val + " " + this.preferredUnits.heartRate;
            },
          },
        },
        {
          labels: {
            formatter: (val) => {
              return val + " " + this.preferredUnits.cadence;
            },
          },
        },
        { show: false },
      ],
    };

    let chart = new ApexCharts(this, options);
    chart.render();
    chart.hideSeries(this.translations.heartrate);
    chart.hideSeries(this.translations.cadence);
  }
}

customElements.define("workout-stats", WorkoutStats);
