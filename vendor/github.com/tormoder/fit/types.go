// Code generated using the program found in 'cmd/fitgen/main.go'. DO NOT EDIT.

// SDK Version: 21.115

package fit

// ActivityClass represents the activity_class FIT type.
type ActivityClass byte

const (
	ActivityClassLevel    ActivityClass = 0x7F // 0 to 100
	ActivityClassLevelMax ActivityClass = 100
	ActivityClassAthlete  ActivityClass = 0x80
	ActivityClassInvalid  ActivityClass = 0xFF
)

// ActivityLevel represents the activity_level FIT type.
type ActivityLevel byte

const (
	ActivityLevelLow     ActivityLevel = 0
	ActivityLevelMedium  ActivityLevel = 1
	ActivityLevelHigh    ActivityLevel = 2
	ActivityLevelInvalid ActivityLevel = 0xFF
)

// ActivityMode represents the activity FIT type.
type ActivityMode byte

const (
	ActivityModeManual         ActivityMode = 0
	ActivityModeAutoMultiSport ActivityMode = 1
	ActivityModeInvalid        ActivityMode = 0xFF
)

// ActivitySubtype represents the activity_subtype FIT type.
type ActivitySubtype byte

const (
	ActivitySubtypeGeneric       ActivitySubtype = 0
	ActivitySubtypeTreadmill     ActivitySubtype = 1  // Run
	ActivitySubtypeStreet        ActivitySubtype = 2  // Run
	ActivitySubtypeTrail         ActivitySubtype = 3  // Run
	ActivitySubtypeTrack         ActivitySubtype = 4  // Run
	ActivitySubtypeSpin          ActivitySubtype = 5  // Cycling
	ActivitySubtypeIndoorCycling ActivitySubtype = 6  // Cycling
	ActivitySubtypeRoad          ActivitySubtype = 7  // Cycling
	ActivitySubtypeMountain      ActivitySubtype = 8  // Cycling
	ActivitySubtypeDownhill      ActivitySubtype = 9  // Cycling
	ActivitySubtypeRecumbent     ActivitySubtype = 10 // Cycling
	ActivitySubtypeCyclocross    ActivitySubtype = 11 // Cycling
	ActivitySubtypeHandCycling   ActivitySubtype = 12 // Cycling
	ActivitySubtypeTrackCycling  ActivitySubtype = 13 // Cycling
	ActivitySubtypeIndoorRowing  ActivitySubtype = 14 // Fitness Equipment
	ActivitySubtypeElliptical    ActivitySubtype = 15 // Fitness Equipment
	ActivitySubtypeStairClimbing ActivitySubtype = 16 // Fitness Equipment
	ActivitySubtypeLapSwimming   ActivitySubtype = 17 // Swimming
	ActivitySubtypeOpenWater     ActivitySubtype = 18 // Swimming
	ActivitySubtypeAll           ActivitySubtype = 254
	ActivitySubtypeInvalid       ActivitySubtype = 0xFF
)

// ActivityType represents the activity_type FIT type.
type ActivityType byte

const (
	ActivityTypeGeneric          ActivityType = 0
	ActivityTypeRunning          ActivityType = 1
	ActivityTypeCycling          ActivityType = 2
	ActivityTypeTransition       ActivityType = 3 // Mulitsport transition
	ActivityTypeFitnessEquipment ActivityType = 4
	ActivityTypeSwimming         ActivityType = 5
	ActivityTypeWalking          ActivityType = 6
	ActivityTypeSedentary        ActivityType = 8
	ActivityTypeAll              ActivityType = 254 // All is for goals only to include all sports.
	ActivityTypeInvalid          ActivityType = 0xFF
)

// AnalogWatchfaceLayout represents the analog_watchface_layout FIT type.
type AnalogWatchfaceLayout byte

const (
	AnalogWatchfaceLayoutMinimal     AnalogWatchfaceLayout = 0
	AnalogWatchfaceLayoutTraditional AnalogWatchfaceLayout = 1
	AnalogWatchfaceLayoutModern      AnalogWatchfaceLayout = 2
	AnalogWatchfaceLayoutInvalid     AnalogWatchfaceLayout = 0xFF
)

// AntChannelId represents the ant_channel_id FIT type.
type AntChannelId uint32

const (
	AntChannelIdAntExtendedDeviceNumberUpperNibble AntChannelId = 0xF0000000
	AntChannelIdAntTransmissionTypeLowerNibble     AntChannelId = 0x0F000000
	AntChannelIdAntDeviceType                      AntChannelId = 0x00FF0000
	AntChannelIdAntDeviceNumber                    AntChannelId = 0x0000FFFF
	AntChannelIdInvalid                            AntChannelId = 0x00000000
)

// AntNetwork represents the ant_network FIT type.
type AntNetwork byte

const (
	AntNetworkPublic  AntNetwork = 0
	AntNetworkAntplus AntNetwork = 1
	AntNetworkAntfs   AntNetwork = 2
	AntNetworkPrivate AntNetwork = 3
	AntNetworkInvalid AntNetwork = 0xFF
)

// AntplusDeviceType represents the antplus_device_type FIT type.
type AntplusDeviceType uint8

const (
	AntplusDeviceTypeAntfs                   AntplusDeviceType = 1
	AntplusDeviceTypeBikePower               AntplusDeviceType = 11
	AntplusDeviceTypeEnvironmentSensorLegacy AntplusDeviceType = 12
	AntplusDeviceTypeMultiSportSpeedDistance AntplusDeviceType = 15
	AntplusDeviceTypeControl                 AntplusDeviceType = 16
	AntplusDeviceTypeFitnessEquipment        AntplusDeviceType = 17
	AntplusDeviceTypeBloodPressure           AntplusDeviceType = 18
	AntplusDeviceTypeGeocacheNode            AntplusDeviceType = 19
	AntplusDeviceTypeLightElectricVehicle    AntplusDeviceType = 20
	AntplusDeviceTypeEnvSensor               AntplusDeviceType = 25
	AntplusDeviceTypeRacquet                 AntplusDeviceType = 26
	AntplusDeviceTypeControlHub              AntplusDeviceType = 27
	AntplusDeviceTypeMuscleOxygen            AntplusDeviceType = 31
	AntplusDeviceTypeShifting                AntplusDeviceType = 34
	AntplusDeviceTypeBikeLightMain           AntplusDeviceType = 35
	AntplusDeviceTypeBikeLightShared         AntplusDeviceType = 36
	AntplusDeviceTypeExd                     AntplusDeviceType = 38
	AntplusDeviceTypeBikeRadar               AntplusDeviceType = 40
	AntplusDeviceTypeBikeAero                AntplusDeviceType = 46
	AntplusDeviceTypeWeightScale             AntplusDeviceType = 119
	AntplusDeviceTypeHeartRate               AntplusDeviceType = 120
	AntplusDeviceTypeBikeSpeedCadence        AntplusDeviceType = 121
	AntplusDeviceTypeBikeCadence             AntplusDeviceType = 122
	AntplusDeviceTypeBikeSpeed               AntplusDeviceType = 123
	AntplusDeviceTypeStrideSpeedDistance     AntplusDeviceType = 124
	AntplusDeviceTypeInvalid                 AntplusDeviceType = 0xFF
)

// AttitudeStage represents the attitude_stage FIT type.
type AttitudeStage byte

const (
	AttitudeStageFailed   AttitudeStage = 0
	AttitudeStageAligning AttitudeStage = 1
	AttitudeStageDegraded AttitudeStage = 2
	AttitudeStageValid    AttitudeStage = 3
	AttitudeStageInvalid  AttitudeStage = 0xFF
)

// AttitudeValidity represents the attitude_validity FIT type.
type AttitudeValidity uint16

const (
	AttitudeValidityTrackAngleHeadingValid AttitudeValidity = 0x0001
	AttitudeValidityPitchValid             AttitudeValidity = 0x0002
	AttitudeValidityRollValid              AttitudeValidity = 0x0004
	AttitudeValidityLateralBodyAccelValid  AttitudeValidity = 0x0008
	AttitudeValidityNormalBodyAccelValid   AttitudeValidity = 0x0010
	AttitudeValidityTurnRateValid          AttitudeValidity = 0x0020
	AttitudeValidityHwFail                 AttitudeValidity = 0x0040
	AttitudeValidityMagInvalid             AttitudeValidity = 0x0080
	AttitudeValidityNoGps                  AttitudeValidity = 0x0100
	AttitudeValidityGpsInvalid             AttitudeValidity = 0x0200
	AttitudeValiditySolutionCoasting       AttitudeValidity = 0x0400
	AttitudeValidityTrueTrackAngle         AttitudeValidity = 0x0800
	AttitudeValidityMagneticHeading        AttitudeValidity = 0x1000
	AttitudeValidityInvalid                AttitudeValidity = 0xFFFF
)

// AutoActivityDetect represents the auto_activity_detect FIT type.
type AutoActivityDetect uint32

const (
	AutoActivityDetectNone       AutoActivityDetect = 0x00000000
	AutoActivityDetectRunning    AutoActivityDetect = 0x00000001
	AutoActivityDetectCycling    AutoActivityDetect = 0x00000002
	AutoActivityDetectSwimming   AutoActivityDetect = 0x00000004
	AutoActivityDetectWalking    AutoActivityDetect = 0x00000008
	AutoActivityDetectElliptical AutoActivityDetect = 0x00000020
	AutoActivityDetectSedentary  AutoActivityDetect = 0x00000400
	AutoActivityDetectInvalid    AutoActivityDetect = 0xFFFFFFFF
)

// AutoSyncFrequency represents the auto_sync_frequency FIT type.
type AutoSyncFrequency byte

const (
	AutoSyncFrequencyNever        AutoSyncFrequency = 0
	AutoSyncFrequencyOccasionally AutoSyncFrequency = 1
	AutoSyncFrequencyFrequent     AutoSyncFrequency = 2
	AutoSyncFrequencyOnceADay     AutoSyncFrequency = 3
	AutoSyncFrequencyRemote       AutoSyncFrequency = 4
	AutoSyncFrequencyInvalid      AutoSyncFrequency = 0xFF
)

// AutolapTrigger represents the autolap_trigger FIT type.
type AutolapTrigger byte

const (
	AutolapTriggerTime             AutolapTrigger = 0
	AutolapTriggerDistance         AutolapTrigger = 1
	AutolapTriggerPositionStart    AutolapTrigger = 2
	AutolapTriggerPositionLap      AutolapTrigger = 3
	AutolapTriggerPositionWaypoint AutolapTrigger = 4
	AutolapTriggerPositionMarked   AutolapTrigger = 5
	AutolapTriggerOff              AutolapTrigger = 6
	AutolapTriggerInvalid          AutolapTrigger = 0xFF
)

// Autoscroll represents the autoscroll FIT type.
type Autoscroll byte

const (
	AutoscrollNone    Autoscroll = 0
	AutoscrollSlow    Autoscroll = 1
	AutoscrollMedium  Autoscroll = 2
	AutoscrollFast    Autoscroll = 3
	AutoscrollInvalid Autoscroll = 0xFF
)

// BacklightMode represents the backlight_mode FIT type.
type BacklightMode byte

const (
	BacklightModeOff                                 BacklightMode = 0
	BacklightModeManual                              BacklightMode = 1
	BacklightModeKeyAndMessages                      BacklightMode = 2
	BacklightModeAutoBrightness                      BacklightMode = 3
	BacklightModeSmartNotifications                  BacklightMode = 4
	BacklightModeKeyAndMessagesNight                 BacklightMode = 5
	BacklightModeKeyAndMessagesAndSmartNotifications BacklightMode = 6
	BacklightModeInvalid                             BacklightMode = 0xFF
)

// BacklightTimeout represents the backlight_timeout FIT type.
type BacklightTimeout uint8

const (
	BacklightTimeoutInfinite BacklightTimeout = 0 // Backlight stays on forever.
	BacklightTimeoutInvalid  BacklightTimeout = 0xFF
)

// BatteryStatus represents the battery_status FIT type.
type BatteryStatus uint8

const (
	BatteryStatusNew      BatteryStatus = 1
	BatteryStatusGood     BatteryStatus = 2
	BatteryStatusOk       BatteryStatus = 3
	BatteryStatusLow      BatteryStatus = 4
	BatteryStatusCritical BatteryStatus = 5
	BatteryStatusCharging BatteryStatus = 6
	BatteryStatusUnknown  BatteryStatus = 7
	BatteryStatusInvalid  BatteryStatus = 0xFF
)

// BenchPressExerciseName represents the bench_press_exercise_name FIT type.
type BenchPressExerciseName uint16

const (
	BenchPressExerciseNameAlternatingDumbbellChestPressOnSwissBall BenchPressExerciseName = 0
	BenchPressExerciseNameBarbellBenchPress                        BenchPressExerciseName = 1
	BenchPressExerciseNameBarbellBoardBenchPress                   BenchPressExerciseName = 2
	BenchPressExerciseNameBarbellFloorPress                        BenchPressExerciseName = 3
	BenchPressExerciseNameCloseGripBarbellBenchPress               BenchPressExerciseName = 4
	BenchPressExerciseNameDeclineDumbbellBenchPress                BenchPressExerciseName = 5
	BenchPressExerciseNameDumbbellBenchPress                       BenchPressExerciseName = 6
	BenchPressExerciseNameDumbbellFloorPress                       BenchPressExerciseName = 7
	BenchPressExerciseNameInclineBarbellBenchPress                 BenchPressExerciseName = 8
	BenchPressExerciseNameInclineDumbbellBenchPress                BenchPressExerciseName = 9
	BenchPressExerciseNameInclineSmithMachineBenchPress            BenchPressExerciseName = 10
	BenchPressExerciseNameIsometricBarbellBenchPress               BenchPressExerciseName = 11
	BenchPressExerciseNameKettlebellChestPress                     BenchPressExerciseName = 12
	BenchPressExerciseNameNeutralGripDumbbellBenchPress            BenchPressExerciseName = 13
	BenchPressExerciseNameNeutralGripDumbbellInclineBenchPress     BenchPressExerciseName = 14
	BenchPressExerciseNameOneArmFloorPress                         BenchPressExerciseName = 15
	BenchPressExerciseNameWeightedOneArmFloorPress                 BenchPressExerciseName = 16
	BenchPressExerciseNamePartialLockout                           BenchPressExerciseName = 17
	BenchPressExerciseNameReverseGripBarbellBenchPress             BenchPressExerciseName = 18
	BenchPressExerciseNameReverseGripInclineBenchPress             BenchPressExerciseName = 19
	BenchPressExerciseNameSingleArmCableChestPress                 BenchPressExerciseName = 20
	BenchPressExerciseNameSingleArmDumbbellBenchPress              BenchPressExerciseName = 21
	BenchPressExerciseNameSmithMachineBenchPress                   BenchPressExerciseName = 22
	BenchPressExerciseNameSwissBallDumbbellChestPress              BenchPressExerciseName = 23
	BenchPressExerciseNameTripleStopBarbellBenchPress              BenchPressExerciseName = 24
	BenchPressExerciseNameWideGripBarbellBenchPress                BenchPressExerciseName = 25
	BenchPressExerciseNameAlternatingDumbbellChestPress            BenchPressExerciseName = 26
	BenchPressExerciseNameInvalid                                  BenchPressExerciseName = 0xFFFF
)

// BikeLightBeamAngleMode represents the bike_light_beam_angle_mode FIT type.
type BikeLightBeamAngleMode uint8

const (
	BikeLightBeamAngleModeManual  BikeLightBeamAngleMode = 0
	BikeLightBeamAngleModeAuto    BikeLightBeamAngleMode = 1
	BikeLightBeamAngleModeInvalid BikeLightBeamAngleMode = 0xFF
)

// BikeLightNetworkConfigType represents the bike_light_network_config_type FIT type.
type BikeLightNetworkConfigType byte

const (
	BikeLightNetworkConfigTypeAuto           BikeLightNetworkConfigType = 0
	BikeLightNetworkConfigTypeIndividual     BikeLightNetworkConfigType = 4
	BikeLightNetworkConfigTypeHighVisibility BikeLightNetworkConfigType = 5
	BikeLightNetworkConfigTypeTrail          BikeLightNetworkConfigType = 6
	BikeLightNetworkConfigTypeInvalid        BikeLightNetworkConfigType = 0xFF
)

// BleDeviceType represents the ble_device_type FIT type.
type BleDeviceType uint8

const (
	BleDeviceTypeConnectedGps     BleDeviceType = 0 // GPS that is provided over a proprietary bluetooth service
	BleDeviceTypeHeartRate        BleDeviceType = 1
	BleDeviceTypeBikePower        BleDeviceType = 2
	BleDeviceTypeBikeSpeedCadence BleDeviceType = 3
	BleDeviceTypeBikeSpeed        BleDeviceType = 4
	BleDeviceTypeBikeCadence      BleDeviceType = 5
	BleDeviceTypeFootpod          BleDeviceType = 6
	BleDeviceTypeBikeTrainer      BleDeviceType = 7 // Indoor-Bike FTMS protocol
	BleDeviceTypeInvalid          BleDeviceType = 0xFF
)

// BodyLocation represents the body_location FIT type.
type BodyLocation byte

const (
	BodyLocationLeftLeg               BodyLocation = 0
	BodyLocationLeftCalf              BodyLocation = 1
	BodyLocationLeftShin              BodyLocation = 2
	BodyLocationLeftHamstring         BodyLocation = 3
	BodyLocationLeftQuad              BodyLocation = 4
	BodyLocationLeftGlute             BodyLocation = 5
	BodyLocationRightLeg              BodyLocation = 6
	BodyLocationRightCalf             BodyLocation = 7
	BodyLocationRightShin             BodyLocation = 8
	BodyLocationRightHamstring        BodyLocation = 9
	BodyLocationRightQuad             BodyLocation = 10
	BodyLocationRightGlute            BodyLocation = 11
	BodyLocationTorsoBack             BodyLocation = 12
	BodyLocationLeftLowerBack         BodyLocation = 13
	BodyLocationLeftUpperBack         BodyLocation = 14
	BodyLocationRightLowerBack        BodyLocation = 15
	BodyLocationRightUpperBack        BodyLocation = 16
	BodyLocationTorsoFront            BodyLocation = 17
	BodyLocationLeftAbdomen           BodyLocation = 18
	BodyLocationLeftChest             BodyLocation = 19
	BodyLocationRightAbdomen          BodyLocation = 20
	BodyLocationRightChest            BodyLocation = 21
	BodyLocationLeftArm               BodyLocation = 22
	BodyLocationLeftShoulder          BodyLocation = 23
	BodyLocationLeftBicep             BodyLocation = 24
	BodyLocationLeftTricep            BodyLocation = 25
	BodyLocationLeftBrachioradialis   BodyLocation = 26 // Left anterior forearm
	BodyLocationLeftForearmExtensors  BodyLocation = 27 // Left posterior forearm
	BodyLocationRightArm              BodyLocation = 28
	BodyLocationRightShoulder         BodyLocation = 29
	BodyLocationRightBicep            BodyLocation = 30
	BodyLocationRightTricep           BodyLocation = 31
	BodyLocationRightBrachioradialis  BodyLocation = 32 // Right anterior forearm
	BodyLocationRightForearmExtensors BodyLocation = 33 // Right posterior forearm
	BodyLocationNeck                  BodyLocation = 34
	BodyLocationThroat                BodyLocation = 35
	BodyLocationWaistMidBack          BodyLocation = 36
	BodyLocationWaistFront            BodyLocation = 37
	BodyLocationWaistLeft             BodyLocation = 38
	BodyLocationWaistRight            BodyLocation = 39
	BodyLocationInvalid               BodyLocation = 0xFF
)

// BpStatus represents the bp_status FIT type.
type BpStatus byte

const (
	BpStatusNoError                 BpStatus = 0
	BpStatusErrorIncompleteData     BpStatus = 1
	BpStatusErrorNoMeasurement      BpStatus = 2
	BpStatusErrorDataOutOfRange     BpStatus = 3
	BpStatusErrorIrregularHeartRate BpStatus = 4
	BpStatusInvalid                 BpStatus = 0xFF
)

// CalfRaiseExerciseName represents the calf_raise_exercise_name FIT type.
type CalfRaiseExerciseName uint16

const (
	CalfRaiseExerciseName3WayCalfRaise                      CalfRaiseExerciseName = 0
	CalfRaiseExerciseName3WayWeightedCalfRaise              CalfRaiseExerciseName = 1
	CalfRaiseExerciseName3WaySingleLegCalfRaise             CalfRaiseExerciseName = 2
	CalfRaiseExerciseName3WayWeightedSingleLegCalfRaise     CalfRaiseExerciseName = 3
	CalfRaiseExerciseNameDonkeyCalfRaise                    CalfRaiseExerciseName = 4
	CalfRaiseExerciseNameWeightedDonkeyCalfRaise            CalfRaiseExerciseName = 5
	CalfRaiseExerciseNameSeatedCalfRaise                    CalfRaiseExerciseName = 6
	CalfRaiseExerciseNameWeightedSeatedCalfRaise            CalfRaiseExerciseName = 7
	CalfRaiseExerciseNameSeatedDumbbellToeRaise             CalfRaiseExerciseName = 8
	CalfRaiseExerciseNameSingleLegBentKneeCalfRaise         CalfRaiseExerciseName = 9
	CalfRaiseExerciseNameWeightedSingleLegBentKneeCalfRaise CalfRaiseExerciseName = 10
	CalfRaiseExerciseNameSingleLegDeclinePushUp             CalfRaiseExerciseName = 11
	CalfRaiseExerciseNameSingleLegDonkeyCalfRaise           CalfRaiseExerciseName = 12
	CalfRaiseExerciseNameWeightedSingleLegDonkeyCalfRaise   CalfRaiseExerciseName = 13
	CalfRaiseExerciseNameSingleLegHipRaiseWithKneeHold      CalfRaiseExerciseName = 14
	CalfRaiseExerciseNameSingleLegStandingCalfRaise         CalfRaiseExerciseName = 15
	CalfRaiseExerciseNameSingleLegStandingDumbbellCalfRaise CalfRaiseExerciseName = 16
	CalfRaiseExerciseNameStandingBarbellCalfRaise           CalfRaiseExerciseName = 17
	CalfRaiseExerciseNameStandingCalfRaise                  CalfRaiseExerciseName = 18
	CalfRaiseExerciseNameWeightedStandingCalfRaise          CalfRaiseExerciseName = 19
	CalfRaiseExerciseNameStandingDumbbellCalfRaise          CalfRaiseExerciseName = 20
	CalfRaiseExerciseNameInvalid                            CalfRaiseExerciseName = 0xFFFF
)

// CameraEventType represents the camera_event_type FIT type.
type CameraEventType byte

const (
	CameraEventTypeVideoStart                  CameraEventType = 0 // Start of video recording
	CameraEventTypeVideoSplit                  CameraEventType = 1 // Mark of video file split (end of one file, beginning of the other)
	CameraEventTypeVideoEnd                    CameraEventType = 2 // End of video recording
	CameraEventTypePhotoTaken                  CameraEventType = 3 // Still photo taken
	CameraEventTypeVideoSecondStreamStart      CameraEventType = 4
	CameraEventTypeVideoSecondStreamSplit      CameraEventType = 5
	CameraEventTypeVideoSecondStreamEnd        CameraEventType = 6
	CameraEventTypeVideoSplitStart             CameraEventType = 7 // Mark of video file split start
	CameraEventTypeVideoSecondStreamSplitStart CameraEventType = 8
	CameraEventTypeVideoPause                  CameraEventType = 11 // Mark when a video recording has been paused
	CameraEventTypeVideoSecondStreamPause      CameraEventType = 12
	CameraEventTypeVideoResume                 CameraEventType = 13 // Mark when a video recording has been resumed
	CameraEventTypeVideoSecondStreamResume     CameraEventType = 14
	CameraEventTypeInvalid                     CameraEventType = 0xFF
)

// CameraOrientationType represents the camera_orientation_type FIT type.
type CameraOrientationType byte

const (
	CameraOrientationTypeCameraOrientation0   CameraOrientationType = 0
	CameraOrientationTypeCameraOrientation90  CameraOrientationType = 1
	CameraOrientationTypeCameraOrientation180 CameraOrientationType = 2
	CameraOrientationTypeCameraOrientation270 CameraOrientationType = 3
	CameraOrientationTypeInvalid              CameraOrientationType = 0xFF
)

// CardioExerciseName represents the cardio_exercise_name FIT type.
type CardioExerciseName uint16

const (
	CardioExerciseNameBobAndWeaveCircle         CardioExerciseName = 0
	CardioExerciseNameWeightedBobAndWeaveCircle CardioExerciseName = 1
	CardioExerciseNameCardioCoreCrawl           CardioExerciseName = 2
	CardioExerciseNameWeightedCardioCoreCrawl   CardioExerciseName = 3
	CardioExerciseNameDoubleUnder               CardioExerciseName = 4
	CardioExerciseNameWeightedDoubleUnder       CardioExerciseName = 5
	CardioExerciseNameJumpRope                  CardioExerciseName = 6
	CardioExerciseNameWeightedJumpRope          CardioExerciseName = 7
	CardioExerciseNameJumpRopeCrossover         CardioExerciseName = 8
	CardioExerciseNameWeightedJumpRopeCrossover CardioExerciseName = 9
	CardioExerciseNameJumpRopeJog               CardioExerciseName = 10
	CardioExerciseNameWeightedJumpRopeJog       CardioExerciseName = 11
	CardioExerciseNameJumpingJacks              CardioExerciseName = 12
	CardioExerciseNameWeightedJumpingJacks      CardioExerciseName = 13
	CardioExerciseNameSkiMoguls                 CardioExerciseName = 14
	CardioExerciseNameWeightedSkiMoguls         CardioExerciseName = 15
	CardioExerciseNameSplitJacks                CardioExerciseName = 16
	CardioExerciseNameWeightedSplitJacks        CardioExerciseName = 17
	CardioExerciseNameSquatJacks                CardioExerciseName = 18
	CardioExerciseNameWeightedSquatJacks        CardioExerciseName = 19
	CardioExerciseNameTripleUnder               CardioExerciseName = 20
	CardioExerciseNameWeightedTripleUnder       CardioExerciseName = 21
	CardioExerciseNameInvalid                   CardioExerciseName = 0xFFFF
)

// CarryExerciseName represents the carry_exercise_name FIT type.
type CarryExerciseName uint16

const (
	CarryExerciseNameBarHolds          CarryExerciseName = 0
	CarryExerciseNameFarmersWalk       CarryExerciseName = 1
	CarryExerciseNameFarmersWalkOnToes CarryExerciseName = 2
	CarryExerciseNameHexDumbbellHold   CarryExerciseName = 3
	CarryExerciseNameOverheadCarry     CarryExerciseName = 4
	CarryExerciseNameInvalid           CarryExerciseName = 0xFFFF
)

// CcrSetpointSwitchMode represents the ccr_setpoint_switch_mode FIT type.
type CcrSetpointSwitchMode byte

const (
	CcrSetpointSwitchModeManual    CcrSetpointSwitchMode = 0 // User switches setpoints manually
	CcrSetpointSwitchModeAutomatic CcrSetpointSwitchMode = 1 // Switch automatically based on depth
	CcrSetpointSwitchModeInvalid   CcrSetpointSwitchMode = 0xFF
)

// Checksum represents the checksum FIT type.
type Checksum uint8

const (
	ChecksumClear   Checksum = 0 // Allows clear of checksum for flash memory where can only write 1 to 0 without erasing sector.
	ChecksumOk      Checksum = 1 // Set to mark checksum as valid if computes to invalid values 0 or 0xFF. Checksum can also be set to ok to save encoding computation time.
	ChecksumInvalid Checksum = 0xFF
)

// ChopExerciseName represents the chop_exercise_name FIT type.
type ChopExerciseName uint16

const (
	ChopExerciseNameCablePullThrough                   ChopExerciseName = 0
	ChopExerciseNameCableRotationalLift                ChopExerciseName = 1
	ChopExerciseNameCableWoodchop                      ChopExerciseName = 2
	ChopExerciseNameCrossChopToKnee                    ChopExerciseName = 3
	ChopExerciseNameWeightedCrossChopToKnee            ChopExerciseName = 4
	ChopExerciseNameDumbbellChop                       ChopExerciseName = 5
	ChopExerciseNameHalfKneelingRotation               ChopExerciseName = 6
	ChopExerciseNameWeightedHalfKneelingRotation       ChopExerciseName = 7
	ChopExerciseNameHalfKneelingRotationalChop         ChopExerciseName = 8
	ChopExerciseNameHalfKneelingRotationalReverseChop  ChopExerciseName = 9
	ChopExerciseNameHalfKneelingStabilityChop          ChopExerciseName = 10
	ChopExerciseNameHalfKneelingStabilityReverseChop   ChopExerciseName = 11
	ChopExerciseNameKneelingRotationalChop             ChopExerciseName = 12
	ChopExerciseNameKneelingRotationalReverseChop      ChopExerciseName = 13
	ChopExerciseNameKneelingStabilityChop              ChopExerciseName = 14
	ChopExerciseNameKneelingWoodchopper                ChopExerciseName = 15
	ChopExerciseNameMedicineBallWoodChops              ChopExerciseName = 16
	ChopExerciseNamePowerSquatChops                    ChopExerciseName = 17
	ChopExerciseNameWeightedPowerSquatChops            ChopExerciseName = 18
	ChopExerciseNameStandingRotationalChop             ChopExerciseName = 19
	ChopExerciseNameStandingSplitRotationalChop        ChopExerciseName = 20
	ChopExerciseNameStandingSplitRotationalReverseChop ChopExerciseName = 21
	ChopExerciseNameStandingStabilityReverseChop       ChopExerciseName = 22
	ChopExerciseNameInvalid                            ChopExerciseName = 0xFFFF
)

// ClimbProEvent represents the climb_pro_event FIT type.
type ClimbProEvent byte

const (
	ClimbProEventApproach ClimbProEvent = 0
	ClimbProEventStart    ClimbProEvent = 1
	ClimbProEventComplete ClimbProEvent = 2
	ClimbProEventInvalid  ClimbProEvent = 0xFF
)

// CommTimeoutType represents the comm_timeout_type FIT type.
type CommTimeoutType uint16

const (
	CommTimeoutTypeWildcardPairingTimeout CommTimeoutType = 0 // Timeout pairing to any device
	CommTimeoutTypePairingTimeout         CommTimeoutType = 1 // Timeout pairing to previously paired device
	CommTimeoutTypeConnectionLost         CommTimeoutType = 2 // Temporary loss of communications
	CommTimeoutTypeConnectionTimeout      CommTimeoutType = 3 // Connection closed due to extended bad communications
	CommTimeoutTypeInvalid                CommTimeoutType = 0xFFFF
)

// ConnectivityCapabilities represents the connectivity_capabilities FIT type.
type ConnectivityCapabilities uint32

const (
	ConnectivityCapabilitiesBluetooth                       ConnectivityCapabilities = 0x00000001
	ConnectivityCapabilitiesBluetoothLe                     ConnectivityCapabilities = 0x00000002
	ConnectivityCapabilitiesAnt                             ConnectivityCapabilities = 0x00000004
	ConnectivityCapabilitiesActivityUpload                  ConnectivityCapabilities = 0x00000008
	ConnectivityCapabilitiesCourseDownload                  ConnectivityCapabilities = 0x00000010
	ConnectivityCapabilitiesWorkoutDownload                 ConnectivityCapabilities = 0x00000020
	ConnectivityCapabilitiesLiveTrack                       ConnectivityCapabilities = 0x00000040
	ConnectivityCapabilitiesWeatherConditions               ConnectivityCapabilities = 0x00000080
	ConnectivityCapabilitiesWeatherAlerts                   ConnectivityCapabilities = 0x00000100
	ConnectivityCapabilitiesGpsEphemerisDownload            ConnectivityCapabilities = 0x00000200
	ConnectivityCapabilitiesExplicitArchive                 ConnectivityCapabilities = 0x00000400
	ConnectivityCapabilitiesSetupIncomplete                 ConnectivityCapabilities = 0x00000800
	ConnectivityCapabilitiesContinueSyncAfterSoftwareUpdate ConnectivityCapabilities = 0x00001000
	ConnectivityCapabilitiesConnectIqAppDownload            ConnectivityCapabilities = 0x00002000
	ConnectivityCapabilitiesGolfCourseDownload              ConnectivityCapabilities = 0x00004000
	ConnectivityCapabilitiesDeviceInitiatesSync             ConnectivityCapabilities = 0x00008000 // Indicates device is in control of initiating all syncs
	ConnectivityCapabilitiesConnectIqWatchAppDownload       ConnectivityCapabilities = 0x00010000
	ConnectivityCapabilitiesConnectIqWidgetDownload         ConnectivityCapabilities = 0x00020000
	ConnectivityCapabilitiesConnectIqWatchFaceDownload      ConnectivityCapabilities = 0x00040000
	ConnectivityCapabilitiesConnectIqDataFieldDownload      ConnectivityCapabilities = 0x00080000
	ConnectivityCapabilitiesConnectIqAppManagment           ConnectivityCapabilities = 0x00100000 // Device supports delete and reorder of apps via GCM
	ConnectivityCapabilitiesSwingSensor                     ConnectivityCapabilities = 0x00200000
	ConnectivityCapabilitiesSwingSensorRemote               ConnectivityCapabilities = 0x00400000
	ConnectivityCapabilitiesIncidentDetection               ConnectivityCapabilities = 0x00800000 // Device supports incident detection
	ConnectivityCapabilitiesAudioPrompts                    ConnectivityCapabilities = 0x01000000
	ConnectivityCapabilitiesWifiVerification                ConnectivityCapabilities = 0x02000000 // Device supports reporting wifi verification via GCM
	ConnectivityCapabilitiesTrueUp                          ConnectivityCapabilities = 0x04000000 // Device supports True Up
	ConnectivityCapabilitiesFindMyWatch                     ConnectivityCapabilities = 0x08000000 // Device supports Find My Watch
	ConnectivityCapabilitiesRemoteManualSync                ConnectivityCapabilities = 0x10000000
	ConnectivityCapabilitiesLiveTrackAutoStart              ConnectivityCapabilities = 0x20000000 // Device supports LiveTrack auto start
	ConnectivityCapabilitiesLiveTrackMessaging              ConnectivityCapabilities = 0x40000000 // Device supports LiveTrack Messaging
	ConnectivityCapabilitiesInstantInput                    ConnectivityCapabilities = 0x80000000 // Device supports instant input feature
	ConnectivityCapabilitiesInvalid                         ConnectivityCapabilities = 0x00000000
)

// CoreExerciseName represents the core_exercise_name FIT type.
type CoreExerciseName uint16

const (
	CoreExerciseNameAbsJabs                          CoreExerciseName = 0
	CoreExerciseNameWeightedAbsJabs                  CoreExerciseName = 1
	CoreExerciseNameAlternatingPlateReach            CoreExerciseName = 2
	CoreExerciseNameBarbellRollout                   CoreExerciseName = 3
	CoreExerciseNameWeightedBarbellRollout           CoreExerciseName = 4
	CoreExerciseNameBodyBarObliqueTwist              CoreExerciseName = 5
	CoreExerciseNameCableCorePress                   CoreExerciseName = 6
	CoreExerciseNameCableSideBend                    CoreExerciseName = 7
	CoreExerciseNameSideBend                         CoreExerciseName = 8
	CoreExerciseNameWeightedSideBend                 CoreExerciseName = 9
	CoreExerciseNameCrescentCircle                   CoreExerciseName = 10
	CoreExerciseNameWeightedCrescentCircle           CoreExerciseName = 11
	CoreExerciseNameCyclingRussianTwist              CoreExerciseName = 12
	CoreExerciseNameWeightedCyclingRussianTwist      CoreExerciseName = 13
	CoreExerciseNameElevatedFeetRussianTwist         CoreExerciseName = 14
	CoreExerciseNameWeightedElevatedFeetRussianTwist CoreExerciseName = 15
	CoreExerciseNameHalfTurkishGetUp                 CoreExerciseName = 16
	CoreExerciseNameKettlebellWindmill               CoreExerciseName = 17
	CoreExerciseNameKneelingAbWheel                  CoreExerciseName = 18
	CoreExerciseNameWeightedKneelingAbWheel          CoreExerciseName = 19
	CoreExerciseNameModifiedFrontLever               CoreExerciseName = 20
	CoreExerciseNameOpenKneeTucks                    CoreExerciseName = 21
	CoreExerciseNameWeightedOpenKneeTucks            CoreExerciseName = 22
	CoreExerciseNameSideAbsLegLift                   CoreExerciseName = 23
	CoreExerciseNameWeightedSideAbsLegLift           CoreExerciseName = 24
	CoreExerciseNameSwissBallJackknife               CoreExerciseName = 25
	CoreExerciseNameWeightedSwissBallJackknife       CoreExerciseName = 26
	CoreExerciseNameSwissBallPike                    CoreExerciseName = 27
	CoreExerciseNameWeightedSwissBallPike            CoreExerciseName = 28
	CoreExerciseNameSwissBallRollout                 CoreExerciseName = 29
	CoreExerciseNameWeightedSwissBallRollout         CoreExerciseName = 30
	CoreExerciseNameTriangleHipPress                 CoreExerciseName = 31
	CoreExerciseNameWeightedTriangleHipPress         CoreExerciseName = 32
	CoreExerciseNameTrxSuspendedJackknife            CoreExerciseName = 33
	CoreExerciseNameWeightedTrxSuspendedJackknife    CoreExerciseName = 34
	CoreExerciseNameUBoat                            CoreExerciseName = 35
	CoreExerciseNameWeightedUBoat                    CoreExerciseName = 36
	CoreExerciseNameWindmillSwitches                 CoreExerciseName = 37
	CoreExerciseNameWeightedWindmillSwitches         CoreExerciseName = 38
	CoreExerciseNameAlternatingSlideOut              CoreExerciseName = 39
	CoreExerciseNameWeightedAlternatingSlideOut      CoreExerciseName = 40
	CoreExerciseNameGhdBackExtensions                CoreExerciseName = 41
	CoreExerciseNameWeightedGhdBackExtensions        CoreExerciseName = 42
	CoreExerciseNameOverheadWalk                     CoreExerciseName = 43
	CoreExerciseNameInchworm                         CoreExerciseName = 44
	CoreExerciseNameWeightedModifiedFrontLever       CoreExerciseName = 45
	CoreExerciseNameRussianTwist                     CoreExerciseName = 46
	CoreExerciseNameAbdominalLegRotations            CoreExerciseName = 47 // Deprecated do not use
	CoreExerciseNameArmAndLegExtensionOnKnees        CoreExerciseName = 48
	CoreExerciseNameBicycle                          CoreExerciseName = 49
	CoreExerciseNameBicepCurlWithLegExtension        CoreExerciseName = 50
	CoreExerciseNameCatCow                           CoreExerciseName = 51
	CoreExerciseNameCorkscrew                        CoreExerciseName = 52
	CoreExerciseNameCrissCross                       CoreExerciseName = 53
	CoreExerciseNameCrissCrossWithBall               CoreExerciseName = 54 // Deprecated do not use
	CoreExerciseNameDoubleLegStretch                 CoreExerciseName = 55
	CoreExerciseNameKneeFolds                        CoreExerciseName = 56
	CoreExerciseNameLowerLift                        CoreExerciseName = 57
	CoreExerciseNameNeckPull                         CoreExerciseName = 58
	CoreExerciseNamePelvicClocks                     CoreExerciseName = 59
	CoreExerciseNameRollOver                         CoreExerciseName = 60
	CoreExerciseNameRollUp                           CoreExerciseName = 61
	CoreExerciseNameRolling                          CoreExerciseName = 62
	CoreExerciseNameRowing1                          CoreExerciseName = 63
	CoreExerciseNameRowing2                          CoreExerciseName = 64
	CoreExerciseNameScissors                         CoreExerciseName = 65
	CoreExerciseNameSingleLegCircles                 CoreExerciseName = 66
	CoreExerciseNameSingleLegStretch                 CoreExerciseName = 67
	CoreExerciseNameSnakeTwist1And2                  CoreExerciseName = 68 // Deprecated do not use
	CoreExerciseNameSwan                             CoreExerciseName = 69
	CoreExerciseNameSwimming                         CoreExerciseName = 70
	CoreExerciseNameTeaser                           CoreExerciseName = 71
	CoreExerciseNameTheHundred                       CoreExerciseName = 72
	CoreExerciseNameInvalid                          CoreExerciseName = 0xFFFF
)

// CourseCapabilities represents the course_capabilities FIT type.
type CourseCapabilities uint32

const (
	CourseCapabilitiesProcessed  CourseCapabilities = 0x00000001
	CourseCapabilitiesValid      CourseCapabilities = 0x00000002
	CourseCapabilitiesTime       CourseCapabilities = 0x00000004
	CourseCapabilitiesDistance   CourseCapabilities = 0x00000008
	CourseCapabilitiesPosition   CourseCapabilities = 0x00000010
	CourseCapabilitiesHeartRate  CourseCapabilities = 0x00000020
	CourseCapabilitiesPower      CourseCapabilities = 0x00000040
	CourseCapabilitiesCadence    CourseCapabilities = 0x00000080
	CourseCapabilitiesTraining   CourseCapabilities = 0x00000100
	CourseCapabilitiesNavigation CourseCapabilities = 0x00000200
	CourseCapabilitiesBikeway    CourseCapabilities = 0x00000400
	CourseCapabilitiesAviation   CourseCapabilities = 0x00001000 // Denote course files to be used as flight plans
	CourseCapabilitiesInvalid    CourseCapabilities = 0x00000000
)

// CoursePoint represents the course_point FIT type.
type CoursePoint byte

const (
	CoursePointGeneric         CoursePoint = 0
	CoursePointSummit          CoursePoint = 1
	CoursePointValley          CoursePoint = 2
	CoursePointWater           CoursePoint = 3
	CoursePointFood            CoursePoint = 4
	CoursePointDanger          CoursePoint = 5
	CoursePointLeft            CoursePoint = 6
	CoursePointRight           CoursePoint = 7
	CoursePointStraight        CoursePoint = 8
	CoursePointFirstAid        CoursePoint = 9
	CoursePointFourthCategory  CoursePoint = 10
	CoursePointThirdCategory   CoursePoint = 11
	CoursePointSecondCategory  CoursePoint = 12
	CoursePointFirstCategory   CoursePoint = 13
	CoursePointHorsCategory    CoursePoint = 14
	CoursePointSprint          CoursePoint = 15
	CoursePointLeftFork        CoursePoint = 16
	CoursePointRightFork       CoursePoint = 17
	CoursePointMiddleFork      CoursePoint = 18
	CoursePointSlightLeft      CoursePoint = 19
	CoursePointSharpLeft       CoursePoint = 20
	CoursePointSlightRight     CoursePoint = 21
	CoursePointSharpRight      CoursePoint = 22
	CoursePointUTurn           CoursePoint = 23
	CoursePointSegmentStart    CoursePoint = 24
	CoursePointSegmentEnd      CoursePoint = 25
	CoursePointCampsite        CoursePoint = 27
	CoursePointAidStation      CoursePoint = 28
	CoursePointRestArea        CoursePoint = 29
	CoursePointGeneralDistance CoursePoint = 30 // Used with UpAhead
	CoursePointService         CoursePoint = 31
	CoursePointEnergyGel       CoursePoint = 32
	CoursePointSportsDrink     CoursePoint = 33
	CoursePointMileMarker      CoursePoint = 34
	CoursePointCheckpoint      CoursePoint = 35
	CoursePointShelter         CoursePoint = 36
	CoursePointMeetingSpot     CoursePoint = 37
	CoursePointOverlook        CoursePoint = 38
	CoursePointToilet          CoursePoint = 39
	CoursePointShower          CoursePoint = 40
	CoursePointGear            CoursePoint = 41
	CoursePointSharpCurve      CoursePoint = 42
	CoursePointSteepIncline    CoursePoint = 43
	CoursePointTunnel          CoursePoint = 44
	CoursePointBridge          CoursePoint = 45
	CoursePointObstacle        CoursePoint = 46
	CoursePointCrossing        CoursePoint = 47
	CoursePointStore           CoursePoint = 48
	CoursePointTransition      CoursePoint = 49
	CoursePointNavaid          CoursePoint = 50
	CoursePointTransport       CoursePoint = 51
	CoursePointAlert           CoursePoint = 52
	CoursePointInfo            CoursePoint = 53
	CoursePointInvalid         CoursePoint = 0xFF
)

// CrunchExerciseName represents the crunch_exercise_name FIT type.
type CrunchExerciseName uint16

const (
	CrunchExerciseNameBicycleCrunch                           CrunchExerciseName = 0
	CrunchExerciseNameCableCrunch                             CrunchExerciseName = 1
	CrunchExerciseNameCircularArmCrunch                       CrunchExerciseName = 2
	CrunchExerciseNameCrossedArmsCrunch                       CrunchExerciseName = 3
	CrunchExerciseNameWeightedCrossedArmsCrunch               CrunchExerciseName = 4
	CrunchExerciseNameCrossLegReverseCrunch                   CrunchExerciseName = 5
	CrunchExerciseNameWeightedCrossLegReverseCrunch           CrunchExerciseName = 6
	CrunchExerciseNameCrunchChop                              CrunchExerciseName = 7
	CrunchExerciseNameWeightedCrunchChop                      CrunchExerciseName = 8
	CrunchExerciseNameDoubleCrunch                            CrunchExerciseName = 9
	CrunchExerciseNameWeightedDoubleCrunch                    CrunchExerciseName = 10
	CrunchExerciseNameElbowToKneeCrunch                       CrunchExerciseName = 11
	CrunchExerciseNameWeightedElbowToKneeCrunch               CrunchExerciseName = 12
	CrunchExerciseNameFlutterKicks                            CrunchExerciseName = 13
	CrunchExerciseNameWeightedFlutterKicks                    CrunchExerciseName = 14
	CrunchExerciseNameFoamRollerReverseCrunchOnBench          CrunchExerciseName = 15
	CrunchExerciseNameWeightedFoamRollerReverseCrunchOnBench  CrunchExerciseName = 16
	CrunchExerciseNameFoamRollerReverseCrunchWithDumbbell     CrunchExerciseName = 17
	CrunchExerciseNameFoamRollerReverseCrunchWithMedicineBall CrunchExerciseName = 18
	CrunchExerciseNameFrogPress                               CrunchExerciseName = 19
	CrunchExerciseNameHangingKneeRaiseObliqueCrunch           CrunchExerciseName = 20
	CrunchExerciseNameWeightedHangingKneeRaiseObliqueCrunch   CrunchExerciseName = 21
	CrunchExerciseNameHipCrossover                            CrunchExerciseName = 22
	CrunchExerciseNameWeightedHipCrossover                    CrunchExerciseName = 23
	CrunchExerciseNameHollowRock                              CrunchExerciseName = 24
	CrunchExerciseNameWeightedHollowRock                      CrunchExerciseName = 25
	CrunchExerciseNameInclineReverseCrunch                    CrunchExerciseName = 26
	CrunchExerciseNameWeightedInclineReverseCrunch            CrunchExerciseName = 27
	CrunchExerciseNameKneelingCableCrunch                     CrunchExerciseName = 28
	CrunchExerciseNameKneelingCrossCrunch                     CrunchExerciseName = 29
	CrunchExerciseNameWeightedKneelingCrossCrunch             CrunchExerciseName = 30
	CrunchExerciseNameKneelingObliqueCableCrunch              CrunchExerciseName = 31
	CrunchExerciseNameKneesToElbow                            CrunchExerciseName = 32
	CrunchExerciseNameLegExtensions                           CrunchExerciseName = 33
	CrunchExerciseNameWeightedLegExtensions                   CrunchExerciseName = 34
	CrunchExerciseNameLegLevers                               CrunchExerciseName = 35
	CrunchExerciseNameMcgillCurlUp                            CrunchExerciseName = 36
	CrunchExerciseNameWeightedMcgillCurlUp                    CrunchExerciseName = 37
	CrunchExerciseNameModifiedPilatesRollUpWithBall           CrunchExerciseName = 38
	CrunchExerciseNameWeightedModifiedPilatesRollUpWithBall   CrunchExerciseName = 39
	CrunchExerciseNamePilatesCrunch                           CrunchExerciseName = 40
	CrunchExerciseNameWeightedPilatesCrunch                   CrunchExerciseName = 41
	CrunchExerciseNamePilatesRollUpWithBall                   CrunchExerciseName = 42
	CrunchExerciseNameWeightedPilatesRollUpWithBall           CrunchExerciseName = 43
	CrunchExerciseNameRaisedLegsCrunch                        CrunchExerciseName = 44
	CrunchExerciseNameWeightedRaisedLegsCrunch                CrunchExerciseName = 45
	CrunchExerciseNameReverseCrunch                           CrunchExerciseName = 46
	CrunchExerciseNameWeightedReverseCrunch                   CrunchExerciseName = 47
	CrunchExerciseNameReverseCrunchOnABench                   CrunchExerciseName = 48
	CrunchExerciseNameWeightedReverseCrunchOnABench           CrunchExerciseName = 49
	CrunchExerciseNameReverseCurlAndLift                      CrunchExerciseName = 50
	CrunchExerciseNameWeightedReverseCurlAndLift              CrunchExerciseName = 51
	CrunchExerciseNameRotationalLift                          CrunchExerciseName = 52
	CrunchExerciseNameWeightedRotationalLift                  CrunchExerciseName = 53
	CrunchExerciseNameSeatedAlternatingReverseCrunch          CrunchExerciseName = 54
	CrunchExerciseNameWeightedSeatedAlternatingReverseCrunch  CrunchExerciseName = 55
	CrunchExerciseNameSeatedLegU                              CrunchExerciseName = 56
	CrunchExerciseNameWeightedSeatedLegU                      CrunchExerciseName = 57
	CrunchExerciseNameSideToSideCrunchAndWeave                CrunchExerciseName = 58
	CrunchExerciseNameWeightedSideToSideCrunchAndWeave        CrunchExerciseName = 59
	CrunchExerciseNameSingleLegReverseCrunch                  CrunchExerciseName = 60
	CrunchExerciseNameWeightedSingleLegReverseCrunch          CrunchExerciseName = 61
	CrunchExerciseNameSkaterCrunchCross                       CrunchExerciseName = 62
	CrunchExerciseNameWeightedSkaterCrunchCross               CrunchExerciseName = 63
	CrunchExerciseNameStandingCableCrunch                     CrunchExerciseName = 64
	CrunchExerciseNameStandingSideCrunch                      CrunchExerciseName = 65
	CrunchExerciseNameStepClimb                               CrunchExerciseName = 66
	CrunchExerciseNameWeightedStepClimb                       CrunchExerciseName = 67
	CrunchExerciseNameSwissBallCrunch                         CrunchExerciseName = 68
	CrunchExerciseNameSwissBallReverseCrunch                  CrunchExerciseName = 69
	CrunchExerciseNameWeightedSwissBallReverseCrunch          CrunchExerciseName = 70
	CrunchExerciseNameSwissBallRussianTwist                   CrunchExerciseName = 71
	CrunchExerciseNameWeightedSwissBallRussianTwist           CrunchExerciseName = 72
	CrunchExerciseNameSwissBallSideCrunch                     CrunchExerciseName = 73
	CrunchExerciseNameWeightedSwissBallSideCrunch             CrunchExerciseName = 74
	CrunchExerciseNameThoracicCrunchesOnFoamRoller            CrunchExerciseName = 75
	CrunchExerciseNameWeightedThoracicCrunchesOnFoamRoller    CrunchExerciseName = 76
	CrunchExerciseNameTricepsCrunch                           CrunchExerciseName = 77
	CrunchExerciseNameWeightedBicycleCrunch                   CrunchExerciseName = 78
	CrunchExerciseNameWeightedCrunch                          CrunchExerciseName = 79
	CrunchExerciseNameWeightedSwissBallCrunch                 CrunchExerciseName = 80
	CrunchExerciseNameToesToBar                               CrunchExerciseName = 81
	CrunchExerciseNameWeightedToesToBar                       CrunchExerciseName = 82
	CrunchExerciseNameCrunch                                  CrunchExerciseName = 83
	CrunchExerciseNameStraightLegCrunchWithBall               CrunchExerciseName = 84
	CrunchExerciseNameInvalid                                 CrunchExerciseName = 0xFFFF
)

// CurlExerciseName represents the curl_exercise_name FIT type.
type CurlExerciseName uint16

const (
	CurlExerciseNameAlternatingDumbbellBicepsCurl             CurlExerciseName = 0
	CurlExerciseNameAlternatingDumbbellBicepsCurlOnSwissBall  CurlExerciseName = 1
	CurlExerciseNameAlternatingInclineDumbbellBicepsCurl      CurlExerciseName = 2
	CurlExerciseNameBarbellBicepsCurl                         CurlExerciseName = 3
	CurlExerciseNameBarbellReverseWristCurl                   CurlExerciseName = 4
	CurlExerciseNameBarbellWristCurl                          CurlExerciseName = 5
	CurlExerciseNameBehindTheBackBarbellReverseWristCurl      CurlExerciseName = 6
	CurlExerciseNameBehindTheBackOneArmCableCurl              CurlExerciseName = 7
	CurlExerciseNameCableBicepsCurl                           CurlExerciseName = 8
	CurlExerciseNameCableHammerCurl                           CurlExerciseName = 9
	CurlExerciseNameCheatingBarbellBicepsCurl                 CurlExerciseName = 10
	CurlExerciseNameCloseGripEzBarBicepsCurl                  CurlExerciseName = 11
	CurlExerciseNameCrossBodyDumbbellHammerCurl               CurlExerciseName = 12
	CurlExerciseNameDeadHangBicepsCurl                        CurlExerciseName = 13
	CurlExerciseNameDeclineHammerCurl                         CurlExerciseName = 14
	CurlExerciseNameDumbbellBicepsCurlWithStaticHold          CurlExerciseName = 15
	CurlExerciseNameDumbbellHammerCurl                        CurlExerciseName = 16
	CurlExerciseNameDumbbellReverseWristCurl                  CurlExerciseName = 17
	CurlExerciseNameDumbbellWristCurl                         CurlExerciseName = 18
	CurlExerciseNameEzBarPreacherCurl                         CurlExerciseName = 19
	CurlExerciseNameForwardBendBicepsCurl                     CurlExerciseName = 20
	CurlExerciseNameHammerCurlToPress                         CurlExerciseName = 21
	CurlExerciseNameInclineDumbbellBicepsCurl                 CurlExerciseName = 22
	CurlExerciseNameInclineOffsetThumbDumbbellCurl            CurlExerciseName = 23
	CurlExerciseNameKettlebellBicepsCurl                      CurlExerciseName = 24
	CurlExerciseNameLyingConcentrationCableCurl               CurlExerciseName = 25
	CurlExerciseNameOneArmPreacherCurl                        CurlExerciseName = 26
	CurlExerciseNamePlatePinchCurl                            CurlExerciseName = 27
	CurlExerciseNamePreacherCurlWithCable                     CurlExerciseName = 28
	CurlExerciseNameReverseEzBarCurl                          CurlExerciseName = 29
	CurlExerciseNameReverseGripWristCurl                      CurlExerciseName = 30
	CurlExerciseNameReverseGripBarbellBicepsCurl              CurlExerciseName = 31
	CurlExerciseNameSeatedAlternatingDumbbellBicepsCurl       CurlExerciseName = 32
	CurlExerciseNameSeatedDumbbellBicepsCurl                  CurlExerciseName = 33
	CurlExerciseNameSeatedReverseDumbbellCurl                 CurlExerciseName = 34
	CurlExerciseNameSplitStanceOffsetPinkyDumbbellCurl        CurlExerciseName = 35
	CurlExerciseNameStandingAlternatingDumbbellCurls          CurlExerciseName = 36
	CurlExerciseNameStandingDumbbellBicepsCurl                CurlExerciseName = 37
	CurlExerciseNameStandingEzBarBicepsCurl                   CurlExerciseName = 38
	CurlExerciseNameStaticCurl                                CurlExerciseName = 39
	CurlExerciseNameSwissBallDumbbellOverheadTricepsExtension CurlExerciseName = 40
	CurlExerciseNameSwissBallEzBarPreacherCurl                CurlExerciseName = 41
	CurlExerciseNameTwistingStandingDumbbellBicepsCurl        CurlExerciseName = 42
	CurlExerciseNameWideGripEzBarBicepsCurl                   CurlExerciseName = 43
	CurlExerciseNameInvalid                                   CurlExerciseName = 0xFFFF
)

// DateMode represents the date_mode FIT type.
type DateMode byte

const (
	DateModeDayMonth DateMode = 0
	DateModeMonthDay DateMode = 1
	DateModeInvalid  DateMode = 0xFF
)

// DayOfWeek represents the day_of_week FIT type.
type DayOfWeek byte

const (
	DayOfWeekSunday    DayOfWeek = 0
	DayOfWeekMonday    DayOfWeek = 1
	DayOfWeekTuesday   DayOfWeek = 2
	DayOfWeekWednesday DayOfWeek = 3
	DayOfWeekThursday  DayOfWeek = 4
	DayOfWeekFriday    DayOfWeek = 5
	DayOfWeekSaturday  DayOfWeek = 6
	DayOfWeekInvalid   DayOfWeek = 0xFF
)

// DeadliftExerciseName represents the deadlift_exercise_name FIT type.
type DeadliftExerciseName uint16

const (
	DeadliftExerciseNameBarbellDeadlift                       DeadliftExerciseName = 0
	DeadliftExerciseNameBarbellStraightLegDeadlift            DeadliftExerciseName = 1
	DeadliftExerciseNameDumbbellDeadlift                      DeadliftExerciseName = 2
	DeadliftExerciseNameDumbbellSingleLegDeadliftToRow        DeadliftExerciseName = 3
	DeadliftExerciseNameDumbbellStraightLegDeadlift           DeadliftExerciseName = 4
	DeadliftExerciseNameKettlebellFloorToShelf                DeadliftExerciseName = 5
	DeadliftExerciseNameOneArmOneLegDeadlift                  DeadliftExerciseName = 6
	DeadliftExerciseNameRackPull                              DeadliftExerciseName = 7
	DeadliftExerciseNameRotationalDumbbellStraightLegDeadlift DeadliftExerciseName = 8
	DeadliftExerciseNameSingleArmDeadlift                     DeadliftExerciseName = 9
	DeadliftExerciseNameSingleLegBarbellDeadlift              DeadliftExerciseName = 10
	DeadliftExerciseNameSingleLegBarbellStraightLegDeadlift   DeadliftExerciseName = 11
	DeadliftExerciseNameSingleLegDeadliftWithBarbell          DeadliftExerciseName = 12
	DeadliftExerciseNameSingleLegRdlCircuit                   DeadliftExerciseName = 13
	DeadliftExerciseNameSingleLegRomanianDeadliftWithDumbbell DeadliftExerciseName = 14
	DeadliftExerciseNameSumoDeadlift                          DeadliftExerciseName = 15
	DeadliftExerciseNameSumoDeadliftHighPull                  DeadliftExerciseName = 16
	DeadliftExerciseNameTrapBarDeadlift                       DeadliftExerciseName = 17
	DeadliftExerciseNameWideGripBarbellDeadlift               DeadliftExerciseName = 18
	DeadliftExerciseNameInvalid                               DeadliftExerciseName = 0xFFFF
)

// DeviceIndex represents the device_index FIT type.
type DeviceIndex uint8

const (
	DeviceIndexCreator DeviceIndex = 0 // Creator of the file is always device index 0.
	DeviceIndexInvalid DeviceIndex = 0xFF
)

// DigitalWatchfaceLayout represents the digital_watchface_layout FIT type.
type DigitalWatchfaceLayout byte

const (
	DigitalWatchfaceLayoutTraditional DigitalWatchfaceLayout = 0
	DigitalWatchfaceLayoutModern      DigitalWatchfaceLayout = 1
	DigitalWatchfaceLayoutBold        DigitalWatchfaceLayout = 2
	DigitalWatchfaceLayoutInvalid     DigitalWatchfaceLayout = 0xFF
)

// DisplayHeart represents the display_heart FIT type.
type DisplayHeart byte

const (
	DisplayHeartBpm     DisplayHeart = 0
	DisplayHeartMax     DisplayHeart = 1
	DisplayHeartReserve DisplayHeart = 2
	DisplayHeartInvalid DisplayHeart = 0xFF
)

// DisplayMeasure represents the display_measure FIT type.
type DisplayMeasure byte

const (
	DisplayMeasureMetric   DisplayMeasure = 0
	DisplayMeasureStatute  DisplayMeasure = 1
	DisplayMeasureNautical DisplayMeasure = 2
	DisplayMeasureInvalid  DisplayMeasure = 0xFF
)

// DisplayOrientation represents the display_orientation FIT type.
type DisplayOrientation byte

const (
	DisplayOrientationAuto             DisplayOrientation = 0 // automatic if the device supports it
	DisplayOrientationPortrait         DisplayOrientation = 1
	DisplayOrientationLandscape        DisplayOrientation = 2
	DisplayOrientationPortraitFlipped  DisplayOrientation = 3 // portrait mode but rotated 180 degrees
	DisplayOrientationLandscapeFlipped DisplayOrientation = 4 // landscape mode but rotated 180 degrees
	DisplayOrientationInvalid          DisplayOrientation = 0xFF
)

// DisplayPosition represents the display_position FIT type.
type DisplayPosition byte

const (
	DisplayPositionDegree               DisplayPosition = 0  // dd.dddddd
	DisplayPositionDegreeMinute         DisplayPosition = 1  // dddmm.mmm
	DisplayPositionDegreeMinuteSecond   DisplayPosition = 2  // dddmmss
	DisplayPositionAustrianGrid         DisplayPosition = 3  // Austrian Grid (BMN)
	DisplayPositionBritishGrid          DisplayPosition = 4  // British National Grid
	DisplayPositionDutchGrid            DisplayPosition = 5  // Dutch grid system
	DisplayPositionHungarianGrid        DisplayPosition = 6  // Hungarian grid system
	DisplayPositionFinnishGrid          DisplayPosition = 7  // Finnish grid system Zone3 KKJ27
	DisplayPositionGermanGrid           DisplayPosition = 8  // Gausss Krueger (German)
	DisplayPositionIcelandicGrid        DisplayPosition = 9  // Icelandic Grid
	DisplayPositionIndonesianEquatorial DisplayPosition = 10 // Indonesian Equatorial LCO
	DisplayPositionIndonesianIrian      DisplayPosition = 11 // Indonesian Irian LCO
	DisplayPositionIndonesianSouthern   DisplayPosition = 12 // Indonesian Southern LCO
	DisplayPositionIndiaZone0           DisplayPosition = 13 // India zone 0
	DisplayPositionIndiaZoneIA          DisplayPosition = 14 // India zone IA
	DisplayPositionIndiaZoneIB          DisplayPosition = 15 // India zone IB
	DisplayPositionIndiaZoneIIA         DisplayPosition = 16 // India zone IIA
	DisplayPositionIndiaZoneIIB         DisplayPosition = 17 // India zone IIB
	DisplayPositionIndiaZoneIIIA        DisplayPosition = 18 // India zone IIIA
	DisplayPositionIndiaZoneIIIB        DisplayPosition = 19 // India zone IIIB
	DisplayPositionIndiaZoneIVA         DisplayPosition = 20 // India zone IVA
	DisplayPositionIndiaZoneIVB         DisplayPosition = 21 // India zone IVB
	DisplayPositionIrishTransverse      DisplayPosition = 22 // Irish Transverse Mercator
	DisplayPositionIrishGrid            DisplayPosition = 23 // Irish Grid
	DisplayPositionLoran                DisplayPosition = 24 // Loran TD
	DisplayPositionMaidenheadGrid       DisplayPosition = 25 // Maidenhead grid system
	DisplayPositionMgrsGrid             DisplayPosition = 26 // MGRS grid system
	DisplayPositionNewZealandGrid       DisplayPosition = 27 // New Zealand grid system
	DisplayPositionNewZealandTransverse DisplayPosition = 28 // New Zealand Transverse Mercator
	DisplayPositionQatarGrid            DisplayPosition = 29 // Qatar National Grid
	DisplayPositionModifiedSwedishGrid  DisplayPosition = 30 // Modified RT-90 (Sweden)
	DisplayPositionSwedishGrid          DisplayPosition = 31 // RT-90 (Sweden)
	DisplayPositionSouthAfricanGrid     DisplayPosition = 32 // South African Grid
	DisplayPositionSwissGrid            DisplayPosition = 33 // Swiss CH-1903 grid
	DisplayPositionTaiwanGrid           DisplayPosition = 34 // Taiwan Grid
	DisplayPositionUnitedStatesGrid     DisplayPosition = 35 // United States National Grid
	DisplayPositionUtmUpsGrid           DisplayPosition = 36 // UTM/UPS grid system
	DisplayPositionWestMalayan          DisplayPosition = 37 // West Malayan RSO
	DisplayPositionBorneoRso            DisplayPosition = 38 // Borneo RSO
	DisplayPositionEstonianGrid         DisplayPosition = 39 // Estonian grid system
	DisplayPositionLatvianGrid          DisplayPosition = 40 // Latvian Transverse Mercator
	DisplayPositionSwedishRef99Grid     DisplayPosition = 41 // Reference Grid 99 TM (Swedish)
	DisplayPositionInvalid              DisplayPosition = 0xFF
)

// DisplayPower represents the display_power FIT type.
type DisplayPower byte

const (
	DisplayPowerWatts      DisplayPower = 0
	DisplayPowerPercentFtp DisplayPower = 1
	DisplayPowerInvalid    DisplayPower = 0xFF
)

// DiveAlarmType represents the dive_alarm_type FIT type.
type DiveAlarmType byte

const (
	DiveAlarmTypeDepth   DiveAlarmType = 0 // Alarm when a certain depth is crossed
	DiveAlarmTypeTime    DiveAlarmType = 1 // Alarm when a certain time has transpired
	DiveAlarmTypeSpeed   DiveAlarmType = 2 // Alarm when a certain ascent or descent rate is exceeded
	DiveAlarmTypeInvalid DiveAlarmType = 0xFF
)

// DiveAlert represents the dive_alert FIT type.
type DiveAlert byte

const (
	DiveAlertNdlReached                DiveAlert = 0
	DiveAlertGasSwitchPrompted         DiveAlert = 1
	DiveAlertNearSurface               DiveAlert = 2
	DiveAlertApproachingNdl            DiveAlert = 3
	DiveAlertPo2Warn                   DiveAlert = 4
	DiveAlertPo2CritHigh               DiveAlert = 5
	DiveAlertPo2CritLow                DiveAlert = 6
	DiveAlertTimeAlert                 DiveAlert = 7
	DiveAlertDepthAlert                DiveAlert = 8
	DiveAlertDecoCeilingBroken         DiveAlert = 9
	DiveAlertDecoComplete              DiveAlert = 10
	DiveAlertSafetyStopBroken          DiveAlert = 11
	DiveAlertSafetyStopComplete        DiveAlert = 12
	DiveAlertCnsWarning                DiveAlert = 13
	DiveAlertCnsCritical               DiveAlert = 14
	DiveAlertOtuWarning                DiveAlert = 15
	DiveAlertOtuCritical               DiveAlert = 16
	DiveAlertAscentCritical            DiveAlert = 17
	DiveAlertAlertDismissedByKey       DiveAlert = 18
	DiveAlertAlertDismissedByTimeout   DiveAlert = 19
	DiveAlertBatteryLow                DiveAlert = 20
	DiveAlertBatteryCritical           DiveAlert = 21
	DiveAlertSafetyStopStarted         DiveAlert = 22
	DiveAlertApproachingFirstDecoStop  DiveAlert = 23
	DiveAlertSetpointSwitchAutoLow     DiveAlert = 24
	DiveAlertSetpointSwitchAutoHigh    DiveAlert = 25
	DiveAlertSetpointSwitchManualLow   DiveAlert = 26
	DiveAlertSetpointSwitchManualHigh  DiveAlert = 27
	DiveAlertAutoSetpointSwitchIgnored DiveAlert = 28
	DiveAlertSwitchedToOpenCircuit     DiveAlert = 29
	DiveAlertSwitchedToClosedCircuit   DiveAlert = 30
	DiveAlertTankBatteryLow            DiveAlert = 32
	DiveAlertPo2CcrDilLow              DiveAlert = 33 // ccr diluent has low po2
	DiveAlertDecoStopCleared           DiveAlert = 34 // a deco stop has been cleared
	DiveAlertApneaNeutralBuoyancy      DiveAlert = 35 // Target Depth Apnea Alarm triggered
	DiveAlertApneaTargetDepth          DiveAlert = 36 // Neutral Buoyance Apnea Alarm triggered
	DiveAlertApneaSurface              DiveAlert = 37 // Surface Apnea Alarm triggered
	DiveAlertApneaHighSpeed            DiveAlert = 38 // High Speed Apnea Alarm triggered
	DiveAlertApneaLowSpeed             DiveAlert = 39 // Low Speed Apnea Alarm triggered
	DiveAlertInvalid                   DiveAlert = 0xFF
)

// DiveBacklightMode represents the dive_backlight_mode FIT type.
type DiveBacklightMode byte

const (
	DiveBacklightModeAtDepth  DiveBacklightMode = 0
	DiveBacklightModeAlwaysOn DiveBacklightMode = 1
	DiveBacklightModeInvalid  DiveBacklightMode = 0xFF
)

// DiveGasMode represents the dive_gas_mode FIT type.
type DiveGasMode byte

const (
	DiveGasModeOpenCircuit          DiveGasMode = 0
	DiveGasModeClosedCircuitDiluent DiveGasMode = 1
	DiveGasModeInvalid              DiveGasMode = 0xFF
)

// DiveGasStatus represents the dive_gas_status FIT type.
type DiveGasStatus byte

const (
	DiveGasStatusDisabled   DiveGasStatus = 0
	DiveGasStatusEnabled    DiveGasStatus = 1
	DiveGasStatusBackupOnly DiveGasStatus = 2
	DiveGasStatusInvalid    DiveGasStatus = 0xFF
)

// Event represents the event FIT type.
type Event byte

const (
	EventTimer                 Event = 0  // Group 0. Start / stop_all
	EventWorkout               Event = 3  // start / stop
	EventWorkoutStep           Event = 4  // Start at beginning of workout. Stop at end of each step.
	EventPowerDown             Event = 5  // stop_all group 0
	EventPowerUp               Event = 6  // stop_all group 0
	EventOffCourse             Event = 7  // start / stop group 0
	EventSession               Event = 8  // Stop at end of each session.
	EventLap                   Event = 9  // Stop at end of each lap.
	EventCoursePoint           Event = 10 // marker
	EventBattery               Event = 11 // marker
	EventVirtualPartnerPace    Event = 12 // Group 1. Start at beginning of activity if VP enabled, when VP pace is changed during activity or VP enabled mid activity. stop_disable when VP disabled.
	EventHrHighAlert           Event = 13 // Group 0. Start / stop when in alert condition.
	EventHrLowAlert            Event = 14 // Group 0. Start / stop when in alert condition.
	EventSpeedHighAlert        Event = 15 // Group 0. Start / stop when in alert condition.
	EventSpeedLowAlert         Event = 16 // Group 0. Start / stop when in alert condition.
	EventCadHighAlert          Event = 17 // Group 0. Start / stop when in alert condition.
	EventCadLowAlert           Event = 18 // Group 0. Start / stop when in alert condition.
	EventPowerHighAlert        Event = 19 // Group 0. Start / stop when in alert condition.
	EventPowerLowAlert         Event = 20 // Group 0. Start / stop when in alert condition.
	EventRecoveryHr            Event = 21 // marker
	EventBatteryLow            Event = 22 // marker
	EventTimeDurationAlert     Event = 23 // Group 1. Start if enabled mid activity (not required at start of activity). Stop when duration is reached. stop_disable if disabled.
	EventDistanceDurationAlert Event = 24 // Group 1. Start if enabled mid activity (not required at start of activity). Stop when duration is reached. stop_disable if disabled.
	EventCalorieDurationAlert  Event = 25 // Group 1. Start if enabled mid activity (not required at start of activity). Stop when duration is reached. stop_disable if disabled.
	EventActivity              Event = 26 // Group 1.. Stop at end of activity.
	EventFitnessEquipment      Event = 27 // marker
	EventLength                Event = 28 // Stop at end of each length.
	EventUserMarker            Event = 32 // marker
	EventSportPoint            Event = 33 // marker
	EventCalibration           Event = 36 // start/stop/marker
	EventFrontGearChange       Event = 42 // marker
	EventRearGearChange        Event = 43 // marker
	EventRiderPositionChange   Event = 44 // marker
	EventElevHighAlert         Event = 45 // Group 0. Start / stop when in alert condition.
	EventElevLowAlert          Event = 46 // Group 0. Start / stop when in alert condition.
	EventCommTimeout           Event = 47 // marker
	EventAutoActivityDetect    Event = 54 // marker
	EventDiveAlert             Event = 56 // marker
	EventDiveGasSwitched       Event = 57 // marker
	EventTankPressureReserve   Event = 71 // marker
	EventTankPressureCritical  Event = 72 // marker
	EventTankLost              Event = 73 // marker
	EventRadarThreatAlert      Event = 75 // start/stop/marker
	EventTankBatteryLow        Event = 76 // marker
	EventTankPodConnected      Event = 81 // marker - tank pod has connected
	EventTankPodDisconnected   Event = 82 // marker - tank pod has lost connection
	EventInvalid               Event = 0xFF
)

// EventType represents the event_type FIT type.
type EventType byte

const (
	EventTypeStart                  EventType = 0
	EventTypeStop                   EventType = 1
	EventTypeConsecutiveDepreciated EventType = 2
	EventTypeMarker                 EventType = 3
	EventTypeStopAll                EventType = 4
	EventTypeBeginDepreciated       EventType = 5
	EventTypeEndDepreciated         EventType = 6
	EventTypeEndAllDepreciated      EventType = 7
	EventTypeStopDisable            EventType = 8
	EventTypeStopDisableAll         EventType = 9
	EventTypeInvalid                EventType = 0xFF
)

// ExdDataUnits represents the exd_data_units FIT type.
type ExdDataUnits byte

const (
	ExdDataUnitsNoUnits                        ExdDataUnits = 0
	ExdDataUnitsLaps                           ExdDataUnits = 1
	ExdDataUnitsMilesPerHour                   ExdDataUnits = 2
	ExdDataUnitsKilometersPerHour              ExdDataUnits = 3
	ExdDataUnitsFeetPerHour                    ExdDataUnits = 4
	ExdDataUnitsMetersPerHour                  ExdDataUnits = 5
	ExdDataUnitsDegreesCelsius                 ExdDataUnits = 6
	ExdDataUnitsDegreesFarenheit               ExdDataUnits = 7
	ExdDataUnitsZone                           ExdDataUnits = 8
	ExdDataUnitsGear                           ExdDataUnits = 9
	ExdDataUnitsRpm                            ExdDataUnits = 10
	ExdDataUnitsBpm                            ExdDataUnits = 11
	ExdDataUnitsDegrees                        ExdDataUnits = 12
	ExdDataUnitsMillimeters                    ExdDataUnits = 13
	ExdDataUnitsMeters                         ExdDataUnits = 14
	ExdDataUnitsKilometers                     ExdDataUnits = 15
	ExdDataUnitsFeet                           ExdDataUnits = 16
	ExdDataUnitsYards                          ExdDataUnits = 17
	ExdDataUnitsKilofeet                       ExdDataUnits = 18
	ExdDataUnitsMiles                          ExdDataUnits = 19
	ExdDataUnitsTime                           ExdDataUnits = 20
	ExdDataUnitsEnumTurnType                   ExdDataUnits = 21
	ExdDataUnitsPercent                        ExdDataUnits = 22
	ExdDataUnitsWatts                          ExdDataUnits = 23
	ExdDataUnitsWattsPerKilogram               ExdDataUnits = 24
	ExdDataUnitsEnumBatteryStatus              ExdDataUnits = 25
	ExdDataUnitsEnumBikeLightBeamAngleMode     ExdDataUnits = 26
	ExdDataUnitsEnumBikeLightBatteryStatus     ExdDataUnits = 27
	ExdDataUnitsEnumBikeLightNetworkConfigType ExdDataUnits = 28
	ExdDataUnitsLights                         ExdDataUnits = 29
	ExdDataUnitsSeconds                        ExdDataUnits = 30
	ExdDataUnitsMinutes                        ExdDataUnits = 31
	ExdDataUnitsHours                          ExdDataUnits = 32
	ExdDataUnitsCalories                       ExdDataUnits = 33
	ExdDataUnitsKilojoules                     ExdDataUnits = 34
	ExdDataUnitsMilliseconds                   ExdDataUnits = 35
	ExdDataUnitsSecondPerMile                  ExdDataUnits = 36
	ExdDataUnitsSecondPerKilometer             ExdDataUnits = 37
	ExdDataUnitsCentimeter                     ExdDataUnits = 38
	ExdDataUnitsEnumCoursePoint                ExdDataUnits = 39
	ExdDataUnitsBradians                       ExdDataUnits = 40
	ExdDataUnitsEnumSport                      ExdDataUnits = 41
	ExdDataUnitsInchesHg                       ExdDataUnits = 42
	ExdDataUnitsMmHg                           ExdDataUnits = 43
	ExdDataUnitsMbars                          ExdDataUnits = 44
	ExdDataUnitsHectoPascals                   ExdDataUnits = 45
	ExdDataUnitsFeetPerMin                     ExdDataUnits = 46
	ExdDataUnitsMetersPerMin                   ExdDataUnits = 47
	ExdDataUnitsMetersPerSec                   ExdDataUnits = 48
	ExdDataUnitsEightCardinal                  ExdDataUnits = 49
	ExdDataUnitsInvalid                        ExdDataUnits = 0xFF
)

// ExdDescriptors represents the exd_descriptors FIT type.
type ExdDescriptors byte

const (
	ExdDescriptorsBikeLightBatteryStatus           ExdDescriptors = 0
	ExdDescriptorsBeamAngleStatus                  ExdDescriptors = 1
	ExdDescriptorsBateryLevel                      ExdDescriptors = 2
	ExdDescriptorsLightNetworkMode                 ExdDescriptors = 3
	ExdDescriptorsNumberLightsConnected            ExdDescriptors = 4
	ExdDescriptorsCadence                          ExdDescriptors = 5
	ExdDescriptorsDistance                         ExdDescriptors = 6
	ExdDescriptorsEstimatedTimeOfArrival           ExdDescriptors = 7
	ExdDescriptorsHeading                          ExdDescriptors = 8
	ExdDescriptorsTime                             ExdDescriptors = 9
	ExdDescriptorsBatteryLevel                     ExdDescriptors = 10
	ExdDescriptorsTrainerResistance                ExdDescriptors = 11
	ExdDescriptorsTrainerTargetPower               ExdDescriptors = 12
	ExdDescriptorsTimeSeated                       ExdDescriptors = 13
	ExdDescriptorsTimeStanding                     ExdDescriptors = 14
	ExdDescriptorsElevation                        ExdDescriptors = 15
	ExdDescriptorsGrade                            ExdDescriptors = 16
	ExdDescriptorsAscent                           ExdDescriptors = 17
	ExdDescriptorsDescent                          ExdDescriptors = 18
	ExdDescriptorsVerticalSpeed                    ExdDescriptors = 19
	ExdDescriptorsDi2BatteryLevel                  ExdDescriptors = 20
	ExdDescriptorsFrontGear                        ExdDescriptors = 21
	ExdDescriptorsRearGear                         ExdDescriptors = 22
	ExdDescriptorsGearRatio                        ExdDescriptors = 23
	ExdDescriptorsHeartRate                        ExdDescriptors = 24
	ExdDescriptorsHeartRateZone                    ExdDescriptors = 25
	ExdDescriptorsTimeInHeartRateZone              ExdDescriptors = 26
	ExdDescriptorsHeartRateReserve                 ExdDescriptors = 27
	ExdDescriptorsCalories                         ExdDescriptors = 28
	ExdDescriptorsGpsAccuracy                      ExdDescriptors = 29
	ExdDescriptorsGpsSignalStrength                ExdDescriptors = 30
	ExdDescriptorsTemperature                      ExdDescriptors = 31
	ExdDescriptorsTimeOfDay                        ExdDescriptors = 32
	ExdDescriptorsBalance                          ExdDescriptors = 33
	ExdDescriptorsPedalSmoothness                  ExdDescriptors = 34
	ExdDescriptorsPower                            ExdDescriptors = 35
	ExdDescriptorsFunctionalThresholdPower         ExdDescriptors = 36
	ExdDescriptorsIntensityFactor                  ExdDescriptors = 37
	ExdDescriptorsWork                             ExdDescriptors = 38
	ExdDescriptorsPowerRatio                       ExdDescriptors = 39
	ExdDescriptorsNormalizedPower                  ExdDescriptors = 40
	ExdDescriptorsTrainingStressScore              ExdDescriptors = 41
	ExdDescriptorsTimeOnZone                       ExdDescriptors = 42
	ExdDescriptorsSpeed                            ExdDescriptors = 43
	ExdDescriptorsLaps                             ExdDescriptors = 44
	ExdDescriptorsReps                             ExdDescriptors = 45
	ExdDescriptorsWorkoutStep                      ExdDescriptors = 46
	ExdDescriptorsCourseDistance                   ExdDescriptors = 47
	ExdDescriptorsNavigationDistance               ExdDescriptors = 48
	ExdDescriptorsCourseEstimatedTimeOfArrival     ExdDescriptors = 49
	ExdDescriptorsNavigationEstimatedTimeOfArrival ExdDescriptors = 50
	ExdDescriptorsCourseTime                       ExdDescriptors = 51
	ExdDescriptorsNavigationTime                   ExdDescriptors = 52
	ExdDescriptorsCourseHeading                    ExdDescriptors = 53
	ExdDescriptorsNavigationHeading                ExdDescriptors = 54
	ExdDescriptorsPowerZone                        ExdDescriptors = 55
	ExdDescriptorsTorqueEffectiveness              ExdDescriptors = 56
	ExdDescriptorsTimerTime                        ExdDescriptors = 57
	ExdDescriptorsPowerWeightRatio                 ExdDescriptors = 58
	ExdDescriptorsLeftPlatformCenterOffset         ExdDescriptors = 59
	ExdDescriptorsRightPlatformCenterOffset        ExdDescriptors = 60
	ExdDescriptorsLeftPowerPhaseStartAngle         ExdDescriptors = 61
	ExdDescriptorsRightPowerPhaseStartAngle        ExdDescriptors = 62
	ExdDescriptorsLeftPowerPhaseFinishAngle        ExdDescriptors = 63
	ExdDescriptorsRightPowerPhaseFinishAngle       ExdDescriptors = 64
	ExdDescriptorsGears                            ExdDescriptors = 65 // Combined gear information
	ExdDescriptorsPace                             ExdDescriptors = 66
	ExdDescriptorsTrainingEffect                   ExdDescriptors = 67
	ExdDescriptorsVerticalOscillation              ExdDescriptors = 68
	ExdDescriptorsVerticalRatio                    ExdDescriptors = 69
	ExdDescriptorsGroundContactTime                ExdDescriptors = 70
	ExdDescriptorsLeftGroundContactTimeBalance     ExdDescriptors = 71
	ExdDescriptorsRightGroundContactTimeBalance    ExdDescriptors = 72
	ExdDescriptorsStrideLength                     ExdDescriptors = 73
	ExdDescriptorsRunningCadence                   ExdDescriptors = 74
	ExdDescriptorsPerformanceCondition             ExdDescriptors = 75
	ExdDescriptorsCourseType                       ExdDescriptors = 76
	ExdDescriptorsTimeInPowerZone                  ExdDescriptors = 77
	ExdDescriptorsNavigationTurn                   ExdDescriptors = 78
	ExdDescriptorsCourseLocation                   ExdDescriptors = 79
	ExdDescriptorsNavigationLocation               ExdDescriptors = 80
	ExdDescriptorsCompass                          ExdDescriptors = 81
	ExdDescriptorsGearCombo                        ExdDescriptors = 82
	ExdDescriptorsMuscleOxygen                     ExdDescriptors = 83
	ExdDescriptorsIcon                             ExdDescriptors = 84
	ExdDescriptorsCompassHeading                   ExdDescriptors = 85
	ExdDescriptorsGpsHeading                       ExdDescriptors = 86
	ExdDescriptorsGpsElevation                     ExdDescriptors = 87
	ExdDescriptorsAnaerobicTrainingEffect          ExdDescriptors = 88
	ExdDescriptorsCourse                           ExdDescriptors = 89
	ExdDescriptorsOffCourse                        ExdDescriptors = 90
	ExdDescriptorsGlideRatio                       ExdDescriptors = 91
	ExdDescriptorsVerticalDistance                 ExdDescriptors = 92
	ExdDescriptorsVmg                              ExdDescriptors = 93
	ExdDescriptorsAmbientPressure                  ExdDescriptors = 94
	ExdDescriptorsPressure                         ExdDescriptors = 95
	ExdDescriptorsVam                              ExdDescriptors = 96
	ExdDescriptorsInvalid                          ExdDescriptors = 0xFF
)

// ExdDisplayType represents the exd_display_type FIT type.
type ExdDisplayType byte

const (
	ExdDisplayTypeNumerical         ExdDisplayType = 0
	ExdDisplayTypeSimple            ExdDisplayType = 1
	ExdDisplayTypeGraph             ExdDisplayType = 2
	ExdDisplayTypeBar               ExdDisplayType = 3
	ExdDisplayTypeCircleGraph       ExdDisplayType = 4
	ExdDisplayTypeVirtualPartner    ExdDisplayType = 5
	ExdDisplayTypeBalance           ExdDisplayType = 6
	ExdDisplayTypeStringList        ExdDisplayType = 7
	ExdDisplayTypeString            ExdDisplayType = 8
	ExdDisplayTypeSimpleDynamicIcon ExdDisplayType = 9
	ExdDisplayTypeGauge             ExdDisplayType = 10
	ExdDisplayTypeInvalid           ExdDisplayType = 0xFF
)

// ExdLayout represents the exd_layout FIT type.
type ExdLayout byte

const (
	ExdLayoutFullScreen                ExdLayout = 0
	ExdLayoutHalfVertical              ExdLayout = 1
	ExdLayoutHalfHorizontal            ExdLayout = 2
	ExdLayoutHalfVerticalRightSplit    ExdLayout = 3
	ExdLayoutHalfHorizontalBottomSplit ExdLayout = 4
	ExdLayoutFullQuarterSplit          ExdLayout = 5
	ExdLayoutHalfVerticalLeftSplit     ExdLayout = 6
	ExdLayoutHalfHorizontalTopSplit    ExdLayout = 7
	ExdLayoutDynamic                   ExdLayout = 8 // The EXD may display the configured concepts in any layout it sees fit.
	ExdLayoutInvalid                   ExdLayout = 0xFF
)

// ExdQualifiers represents the exd_qualifiers FIT type.
type ExdQualifiers byte

const (
	ExdQualifiersNoQualifier              ExdQualifiers = 0
	ExdQualifiersInstantaneous            ExdQualifiers = 1
	ExdQualifiersAverage                  ExdQualifiers = 2
	ExdQualifiersLap                      ExdQualifiers = 3
	ExdQualifiersMaximum                  ExdQualifiers = 4
	ExdQualifiersMaximumAverage           ExdQualifiers = 5
	ExdQualifiersMaximumLap               ExdQualifiers = 6
	ExdQualifiersLastLap                  ExdQualifiers = 7
	ExdQualifiersAverageLap               ExdQualifiers = 8
	ExdQualifiersToDestination            ExdQualifiers = 9
	ExdQualifiersToGo                     ExdQualifiers = 10
	ExdQualifiersToNext                   ExdQualifiers = 11
	ExdQualifiersNextCoursePoint          ExdQualifiers = 12
	ExdQualifiersTotal                    ExdQualifiers = 13
	ExdQualifiersThreeSecondAverage       ExdQualifiers = 14
	ExdQualifiersTenSecondAverage         ExdQualifiers = 15
	ExdQualifiersThirtySecondAverage      ExdQualifiers = 16
	ExdQualifiersPercentMaximum           ExdQualifiers = 17
	ExdQualifiersPercentMaximumAverage    ExdQualifiers = 18
	ExdQualifiersLapPercentMaximum        ExdQualifiers = 19
	ExdQualifiersElapsed                  ExdQualifiers = 20
	ExdQualifiersSunrise                  ExdQualifiers = 21
	ExdQualifiersSunset                   ExdQualifiers = 22
	ExdQualifiersComparedToVirtualPartner ExdQualifiers = 23
	ExdQualifiersMaximum24h               ExdQualifiers = 24
	ExdQualifiersMinimum24h               ExdQualifiers = 25
	ExdQualifiersMinimum                  ExdQualifiers = 26
	ExdQualifiersFirst                    ExdQualifiers = 27
	ExdQualifiersSecond                   ExdQualifiers = 28
	ExdQualifiersThird                    ExdQualifiers = 29
	ExdQualifiersShifter                  ExdQualifiers = 30
	ExdQualifiersLastSport                ExdQualifiers = 31
	ExdQualifiersMoving                   ExdQualifiers = 32
	ExdQualifiersStopped                  ExdQualifiers = 33
	ExdQualifiersEstimatedTotal           ExdQualifiers = 34
	ExdQualifiersZone9                    ExdQualifiers = 242
	ExdQualifiersZone8                    ExdQualifiers = 243
	ExdQualifiersZone7                    ExdQualifiers = 244
	ExdQualifiersZone6                    ExdQualifiers = 245
	ExdQualifiersZone5                    ExdQualifiers = 246
	ExdQualifiersZone4                    ExdQualifiers = 247
	ExdQualifiersZone3                    ExdQualifiers = 248
	ExdQualifiersZone2                    ExdQualifiers = 249
	ExdQualifiersZone1                    ExdQualifiers = 250
	ExdQualifiersInvalid                  ExdQualifiers = 0xFF
)

// ExerciseCategory represents the exercise_category FIT type.
type ExerciseCategory uint16

const (
	ExerciseCategoryBenchPress        ExerciseCategory = 0
	ExerciseCategoryCalfRaise         ExerciseCategory = 1
	ExerciseCategoryCardio            ExerciseCategory = 2
	ExerciseCategoryCarry             ExerciseCategory = 3
	ExerciseCategoryChop              ExerciseCategory = 4
	ExerciseCategoryCore              ExerciseCategory = 5
	ExerciseCategoryCrunch            ExerciseCategory = 6
	ExerciseCategoryCurl              ExerciseCategory = 7
	ExerciseCategoryDeadlift          ExerciseCategory = 8
	ExerciseCategoryFlye              ExerciseCategory = 9
	ExerciseCategoryHipRaise          ExerciseCategory = 10
	ExerciseCategoryHipStability      ExerciseCategory = 11
	ExerciseCategoryHipSwing          ExerciseCategory = 12
	ExerciseCategoryHyperextension    ExerciseCategory = 13
	ExerciseCategoryLateralRaise      ExerciseCategory = 14
	ExerciseCategoryLegCurl           ExerciseCategory = 15
	ExerciseCategoryLegRaise          ExerciseCategory = 16
	ExerciseCategoryLunge             ExerciseCategory = 17
	ExerciseCategoryOlympicLift       ExerciseCategory = 18
	ExerciseCategoryPlank             ExerciseCategory = 19
	ExerciseCategoryPlyo              ExerciseCategory = 20
	ExerciseCategoryPullUp            ExerciseCategory = 21
	ExerciseCategoryPushUp            ExerciseCategory = 22
	ExerciseCategoryRow               ExerciseCategory = 23
	ExerciseCategoryShoulderPress     ExerciseCategory = 24
	ExerciseCategoryShoulderStability ExerciseCategory = 25
	ExerciseCategoryShrug             ExerciseCategory = 26
	ExerciseCategorySitUp             ExerciseCategory = 27
	ExerciseCategorySquat             ExerciseCategory = 28
	ExerciseCategoryTotalBody         ExerciseCategory = 29
	ExerciseCategoryTricepsExtension  ExerciseCategory = 30
	ExerciseCategoryWarmUp            ExerciseCategory = 31
	ExerciseCategoryRun               ExerciseCategory = 32
	ExerciseCategoryUnknown           ExerciseCategory = 65534
	ExerciseCategoryInvalid           ExerciseCategory = 0xFFFF
)

// FaveroProduct represents the favero_product FIT type.
type FaveroProduct uint16

const (
	FaveroProductAssiomaUno FaveroProduct = 10
	FaveroProductAssiomaDuo FaveroProduct = 12
	FaveroProductInvalid    FaveroProduct = 0xFFFF
)

// FileFlags represents the file_flags FIT type.
type FileFlags uint8

const (
	FileFlagsRead    FileFlags = 0x02
	FileFlagsWrite   FileFlags = 0x04
	FileFlagsErase   FileFlags = 0x08
	FileFlagsInvalid FileFlags = 0x00
)

// FileType represents the file FIT type.
type FileType byte

const (
	FileTypeDevice           FileType = 1  // Read only, single file. Must be in root directory.
	FileTypeSettings         FileType = 2  // Read/write, single file. Directory=Settings
	FileTypeSport            FileType = 3  // Read/write, multiple files, file number = sport type. Directory=Sports
	FileTypeActivity         FileType = 4  // Read/erase, multiple files. Directory=Activities
	FileTypeWorkout          FileType = 5  // Read/write/erase, multiple files. Directory=Workouts
	FileTypeCourse           FileType = 6  // Read/write/erase, multiple files. Directory=Courses
	FileTypeSchedules        FileType = 7  // Read/write, single file. Directory=Schedules
	FileTypeWeight           FileType = 9  // Read only, single file. Circular buffer. All message definitions at start of file. Directory=Weight
	FileTypeTotals           FileType = 10 // Read only, single file. Directory=Totals
	FileTypeGoals            FileType = 11 // Read/write, single file. Directory=Goals
	FileTypeBloodPressure    FileType = 14 // Read only. Directory=Blood Pressure
	FileTypeMonitoringA      FileType = 15 // Read only. Directory=Monitoring. File number=sub type.
	FileTypeActivitySummary  FileType = 20 // Read/erase, multiple files. Directory=Activities
	FileTypeMonitoringDaily  FileType = 28
	FileTypeMonitoringB      FileType = 32   // Read only. Directory=Monitoring. File number=identifier
	FileTypeSegment          FileType = 34   // Read/write/erase. Multiple Files. Directory=Segments
	FileTypeSegmentList      FileType = 35   // Read/write/erase. Single File. Directory=Segments
	FileTypeExdConfiguration FileType = 40   // Read/write/erase. Single File. Directory=Settings
	FileTypeMfgRangeMin      FileType = 0xF7 // 0xF7 - 0xFE reserved for manufacturer specific file types
	FileTypeMfgRangeMax      FileType = 0xFE // 0xF7 - 0xFE reserved for manufacturer specific file types
	FileTypeInvalid          FileType = 0xFF
)

// FitBaseType represents the fit_base_type FIT type.
type FitBaseType uint8

const (
	FitBaseTypeEnum    FitBaseType = 0
	FitBaseTypeSint8   FitBaseType = 1
	FitBaseTypeUint8   FitBaseType = 2
	FitBaseTypeSint16  FitBaseType = 131
	FitBaseTypeUint16  FitBaseType = 132
	FitBaseTypeSint32  FitBaseType = 133
	FitBaseTypeUint32  FitBaseType = 134
	FitBaseTypeString  FitBaseType = 7
	FitBaseTypeFloat32 FitBaseType = 136
	FitBaseTypeFloat64 FitBaseType = 137
	FitBaseTypeUint8z  FitBaseType = 10
	FitBaseTypeUint16z FitBaseType = 139
	FitBaseTypeUint32z FitBaseType = 140
	FitBaseTypeByte    FitBaseType = 13
	FitBaseTypeSint64  FitBaseType = 142
	FitBaseTypeUint64  FitBaseType = 143
	FitBaseTypeUint64z FitBaseType = 144
	FitBaseTypeInvalid FitBaseType = 0xFF
)

// FitBaseUnit represents the fit_base_unit FIT type.
type FitBaseUnit uint16

const (
	FitBaseUnitOther    FitBaseUnit = 0
	FitBaseUnitKilogram FitBaseUnit = 1
	FitBaseUnitPound    FitBaseUnit = 2
	FitBaseUnitInvalid  FitBaseUnit = 0xFFFF
)

// FitnessEquipmentState represents the fitness_equipment_state FIT type.
type FitnessEquipmentState byte

const (
	FitnessEquipmentStateReady   FitnessEquipmentState = 0
	FitnessEquipmentStateInUse   FitnessEquipmentState = 1
	FitnessEquipmentStatePaused  FitnessEquipmentState = 2
	FitnessEquipmentStateUnknown FitnessEquipmentState = 3 // lost connection to fitness equipment
	FitnessEquipmentStateInvalid FitnessEquipmentState = 0xFF
)

// FlyeExerciseName represents the flye_exercise_name FIT type.
type FlyeExerciseName uint16

const (
	FlyeExerciseNameCableCrossover                    FlyeExerciseName = 0
	FlyeExerciseNameDeclineDumbbellFlye               FlyeExerciseName = 1
	FlyeExerciseNameDumbbellFlye                      FlyeExerciseName = 2
	FlyeExerciseNameInclineDumbbellFlye               FlyeExerciseName = 3
	FlyeExerciseNameKettlebellFlye                    FlyeExerciseName = 4
	FlyeExerciseNameKneelingRearFlye                  FlyeExerciseName = 5
	FlyeExerciseNameSingleArmStandingCableReverseFlye FlyeExerciseName = 6
	FlyeExerciseNameSwissBallDumbbellFlye             FlyeExerciseName = 7
	FlyeExerciseNameArmRotations                      FlyeExerciseName = 8
	FlyeExerciseNameHugATree                          FlyeExerciseName = 9
	FlyeExerciseNameInvalid                           FlyeExerciseName = 0xFFFF
)

// GarminProduct represents the garmin_product FIT type.
type GarminProduct uint16

const (
	GarminProductHrm1                       GarminProduct = 1
	GarminProductAxh01                      GarminProduct = 2 // AXH01 HRM chipset
	GarminProductAxb01                      GarminProduct = 3
	GarminProductAxb02                      GarminProduct = 4
	GarminProductHrm2ss                     GarminProduct = 5
	GarminProductDsiAlf02                   GarminProduct = 6
	GarminProductHrm3ss                     GarminProduct = 7
	GarminProductHrmRunSingleByteProductId  GarminProduct = 8   // hrm_run model for HRM ANT+ messaging
	GarminProductBsm                        GarminProduct = 9   // BSM model for ANT+ messaging
	GarminProductBcm                        GarminProduct = 10  // BCM model for ANT+ messaging
	GarminProductAxs01                      GarminProduct = 11  // AXS01 HRM Bike Chipset model for ANT+ messaging
	GarminProductHrmTriSingleByteProductId  GarminProduct = 12  // hrm_tri model for HRM ANT+ messaging
	GarminProductHrm4RunSingleByteProductId GarminProduct = 13  // hrm4 run model for HRM ANT+ messaging
	GarminProductFr225SingleByteProductId   GarminProduct = 14  // fr225 model for HRM ANT+ messaging
	GarminProductGen3BsmSingleByteProductId GarminProduct = 15  // gen3_bsm model for Bike Speed ANT+ messaging
	GarminProductGen3BcmSingleByteProductId GarminProduct = 16  // gen3_bcm model for Bike Cadence ANT+ messaging
	GarminProductOHR                        GarminProduct = 255 // Garmin Wearable Optical Heart Rate Sensor for ANT+ HR Profile Broadcasting
	GarminProductFr301China                 GarminProduct = 473
	GarminProductFr301Japan                 GarminProduct = 474
	GarminProductFr301Korea                 GarminProduct = 475
	GarminProductFr301Taiwan                GarminProduct = 494
	GarminProductFr405                      GarminProduct = 717 // Forerunner 405
	GarminProductFr50                       GarminProduct = 782 // Forerunner 50
	GarminProductFr405Japan                 GarminProduct = 987
	GarminProductFr60                       GarminProduct = 988 // Forerunner 60
	GarminProductDsiAlf01                   GarminProduct = 1011
	GarminProductFr310xt                    GarminProduct = 1018 // Forerunner 310
	GarminProductEdge500                    GarminProduct = 1036
	GarminProductFr110                      GarminProduct = 1124 // Forerunner 110
	GarminProductEdge800                    GarminProduct = 1169
	GarminProductEdge500Taiwan              GarminProduct = 1199
	GarminProductEdge500Japan               GarminProduct = 1213
	GarminProductChirp                      GarminProduct = 1253
	GarminProductFr110Japan                 GarminProduct = 1274
	GarminProductEdge200                    GarminProduct = 1325
	GarminProductFr910xt                    GarminProduct = 1328
	GarminProductEdge800Taiwan              GarminProduct = 1333
	GarminProductEdge800Japan               GarminProduct = 1334
	GarminProductAlf04                      GarminProduct = 1341
	GarminProductFr610                      GarminProduct = 1345
	GarminProductFr210Japan                 GarminProduct = 1360
	GarminProductVectorSs                   GarminProduct = 1380
	GarminProductVectorCp                   GarminProduct = 1381
	GarminProductEdge800China               GarminProduct = 1386
	GarminProductEdge500China               GarminProduct = 1387
	GarminProductApproachG10                GarminProduct = 1405
	GarminProductFr610Japan                 GarminProduct = 1410
	GarminProductEdge500Korea               GarminProduct = 1422
	GarminProductFr70                       GarminProduct = 1436
	GarminProductFr310xt4t                  GarminProduct = 1446
	GarminProductAmx                        GarminProduct = 1461
	GarminProductFr10                       GarminProduct = 1482
	GarminProductEdge800Korea               GarminProduct = 1497
	GarminProductSwim                       GarminProduct = 1499
	GarminProductFr910xtChina               GarminProduct = 1537
	GarminProductFenix                      GarminProduct = 1551
	GarminProductEdge200Taiwan              GarminProduct = 1555
	GarminProductEdge510                    GarminProduct = 1561
	GarminProductEdge810                    GarminProduct = 1567
	GarminProductTempe                      GarminProduct = 1570
	GarminProductFr910xtJapan               GarminProduct = 1600
	GarminProductFr620                      GarminProduct = 1623
	GarminProductFr220                      GarminProduct = 1632
	GarminProductFr910xtKorea               GarminProduct = 1664
	GarminProductFr10Japan                  GarminProduct = 1688
	GarminProductEdge810Japan               GarminProduct = 1721
	GarminProductVirbElite                  GarminProduct = 1735
	GarminProductEdgeTouring                GarminProduct = 1736 // Also Edge Touring Plus
	GarminProductEdge510Japan               GarminProduct = 1742
	GarminProductHrmTri                     GarminProduct = 1743 // Also HRM-Swim
	GarminProductHrmRun                     GarminProduct = 1752
	GarminProductFr920xt                    GarminProduct = 1765
	GarminProductEdge510Asia                GarminProduct = 1821
	GarminProductEdge810China               GarminProduct = 1822
	GarminProductEdge810Taiwan              GarminProduct = 1823
	GarminProductEdge1000                   GarminProduct = 1836
	GarminProductVivoFit                    GarminProduct = 1837
	GarminProductVirbRemote                 GarminProduct = 1853
	GarminProductVivoKi                     GarminProduct = 1885
	GarminProductFr15                       GarminProduct = 1903
	GarminProductVivoActive                 GarminProduct = 1907
	GarminProductEdge510Korea               GarminProduct = 1918
	GarminProductFr620Japan                 GarminProduct = 1928
	GarminProductFr620China                 GarminProduct = 1929
	GarminProductFr220Japan                 GarminProduct = 1930
	GarminProductFr220China                 GarminProduct = 1931
	GarminProductApproachS6                 GarminProduct = 1936
	GarminProductVivoSmart                  GarminProduct = 1956
	GarminProductFenix2                     GarminProduct = 1967
	GarminProductEpix                       GarminProduct = 1988
	GarminProductFenix3                     GarminProduct = 2050
	GarminProductEdge1000Taiwan             GarminProduct = 2052
	GarminProductEdge1000Japan              GarminProduct = 2053
	GarminProductFr15Japan                  GarminProduct = 2061
	GarminProductEdge520                    GarminProduct = 2067
	GarminProductEdge1000China              GarminProduct = 2070
	GarminProductFr620Russia                GarminProduct = 2072
	GarminProductFr220Russia                GarminProduct = 2073
	GarminProductVectorS                    GarminProduct = 2079
	GarminProductEdge1000Korea              GarminProduct = 2100
	GarminProductFr920xtTaiwan              GarminProduct = 2130
	GarminProductFr920xtChina               GarminProduct = 2131
	GarminProductFr920xtJapan               GarminProduct = 2132
	GarminProductVirbx                      GarminProduct = 2134
	GarminProductVivoSmartApac              GarminProduct = 2135
	GarminProductEtrexTouch                 GarminProduct = 2140
	GarminProductEdge25                     GarminProduct = 2147
	GarminProductFr25                       GarminProduct = 2148
	GarminProductVivoFit2                   GarminProduct = 2150
	GarminProductFr225                      GarminProduct = 2153
	GarminProductFr630                      GarminProduct = 2156
	GarminProductFr230                      GarminProduct = 2157
	GarminProductFr735xt                    GarminProduct = 2158
	GarminProductVivoActiveApac             GarminProduct = 2160
	GarminProductVector2                    GarminProduct = 2161
	GarminProductVector2s                   GarminProduct = 2162
	GarminProductVirbxe                     GarminProduct = 2172
	GarminProductFr620Taiwan                GarminProduct = 2173
	GarminProductFr220Taiwan                GarminProduct = 2174
	GarminProductTruswing                   GarminProduct = 2175
	GarminProductD2airvenu                  GarminProduct = 2187
	GarminProductFenix3China                GarminProduct = 2188
	GarminProductFenix3Twn                  GarminProduct = 2189
	GarminProductVariaHeadlight             GarminProduct = 2192
	GarminProductVariaTaillightOld          GarminProduct = 2193
	GarminProductEdgeExplore1000            GarminProduct = 2204
	GarminProductFr225Asia                  GarminProduct = 2219
	GarminProductVariaRadarTaillight        GarminProduct = 2225
	GarminProductVariaRadarDisplay          GarminProduct = 2226
	GarminProductEdge20                     GarminProduct = 2238
	GarminProductEdge520Asia                GarminProduct = 2260
	GarminProductEdge520Japan               GarminProduct = 2261
	GarminProductD2Bravo                    GarminProduct = 2262
	GarminProductApproachS20                GarminProduct = 2266
	GarminProductVivoSmart2                 GarminProduct = 2271
	GarminProductEdge1000Thai               GarminProduct = 2274
	GarminProductVariaRemote                GarminProduct = 2276
	GarminProductEdge25Asia                 GarminProduct = 2288
	GarminProductEdge25Jpn                  GarminProduct = 2289
	GarminProductEdge20Asia                 GarminProduct = 2290
	GarminProductApproachX40                GarminProduct = 2292
	GarminProductFenix3Japan                GarminProduct = 2293
	GarminProductVivoSmartEmea              GarminProduct = 2294
	GarminProductFr630Asia                  GarminProduct = 2310
	GarminProductFr630Jpn                   GarminProduct = 2311
	GarminProductFr230Jpn                   GarminProduct = 2313
	GarminProductHrm4Run                    GarminProduct = 2327
	GarminProductEpixJapan                  GarminProduct = 2332
	GarminProductVivoActiveHr               GarminProduct = 2337
	GarminProductVivoSmartGpsHr             GarminProduct = 2347
	GarminProductVivoSmartHr                GarminProduct = 2348
	GarminProductVivoSmartHrAsia            GarminProduct = 2361
	GarminProductVivoSmartGpsHrAsia         GarminProduct = 2362
	GarminProductVivoMove                   GarminProduct = 2368
	GarminProductVariaTaillight             GarminProduct = 2379
	GarminProductFr235Asia                  GarminProduct = 2396
	GarminProductFr235Japan                 GarminProduct = 2397
	GarminProductVariaVision                GarminProduct = 2398
	GarminProductVivoFit3                   GarminProduct = 2406
	GarminProductFenix3Korea                GarminProduct = 2407
	GarminProductFenix3Sea                  GarminProduct = 2408
	GarminProductFenix3Hr                   GarminProduct = 2413
	GarminProductVirbUltra30                GarminProduct = 2417
	GarminProductIndexSmartScale            GarminProduct = 2429
	GarminProductFr235                      GarminProduct = 2431
	GarminProductFenix3Chronos              GarminProduct = 2432
	GarminProductOregon7xx                  GarminProduct = 2441
	GarminProductRino7xx                    GarminProduct = 2444
	GarminProductEpixKorea                  GarminProduct = 2457
	GarminProductFenix3HrChn                GarminProduct = 2473
	GarminProductFenix3HrTwn                GarminProduct = 2474
	GarminProductFenix3HrJpn                GarminProduct = 2475
	GarminProductFenix3HrSea                GarminProduct = 2476
	GarminProductFenix3HrKor                GarminProduct = 2477
	GarminProductNautix                     GarminProduct = 2496
	GarminProductVivoActiveHrApac           GarminProduct = 2497
	GarminProductFr35                       GarminProduct = 2503
	GarminProductOregon7xxWw                GarminProduct = 2512
	GarminProductEdge820                    GarminProduct = 2530
	GarminProductEdgeExplore820             GarminProduct = 2531
	GarminProductFr735xtApac                GarminProduct = 2533
	GarminProductFr735xtJapan               GarminProduct = 2534
	GarminProductFenix5s                    GarminProduct = 2544
	GarminProductD2BravoTitanium            GarminProduct = 2547
	GarminProductVariaUt800                 GarminProduct = 2567 // Varia UT 800 SW
	GarminProductRunningDynamicsPod         GarminProduct = 2593
	GarminProductEdge820China               GarminProduct = 2599
	GarminProductEdge820Japan               GarminProduct = 2600
	GarminProductFenix5x                    GarminProduct = 2604
	GarminProductVivoFitJr                  GarminProduct = 2606
	GarminProductVivoSmart3                 GarminProduct = 2622
	GarminProductVivoSport                  GarminProduct = 2623
	GarminProductEdge820Taiwan              GarminProduct = 2628
	GarminProductEdge820Korea               GarminProduct = 2629
	GarminProductEdge820Sea                 GarminProduct = 2630
	GarminProductFr35Hebrew                 GarminProduct = 2650
	GarminProductApproachS60                GarminProduct = 2656
	GarminProductFr35Apac                   GarminProduct = 2667
	GarminProductFr35Japan                  GarminProduct = 2668
	GarminProductFenix3ChronosAsia          GarminProduct = 2675
	GarminProductVirb360                    GarminProduct = 2687
	GarminProductFr935                      GarminProduct = 2691
	GarminProductFenix5                     GarminProduct = 2697
	GarminProductVivoactive3                GarminProduct = 2700
	GarminProductFr235ChinaNfc              GarminProduct = 2733
	GarminProductForetrex601701             GarminProduct = 2769
	GarminProductVivoMoveHr                 GarminProduct = 2772
	GarminProductEdge1030                   GarminProduct = 2713
	GarminProductFr35Sea                    GarminProduct = 2727
	GarminProductVector3                    GarminProduct = 2787
	GarminProductFenix5Asia                 GarminProduct = 2796
	GarminProductFenix5sAsia                GarminProduct = 2797
	GarminProductFenix5xAsia                GarminProduct = 2798
	GarminProductApproachZ80                GarminProduct = 2806
	GarminProductFr35Korea                  GarminProduct = 2814
	GarminProductD2charlie                  GarminProduct = 2819
	GarminProductVivoSmart3Apac             GarminProduct = 2831
	GarminProductVivoSportApac              GarminProduct = 2832
	GarminProductFr935Asia                  GarminProduct = 2833
	GarminProductDescent                    GarminProduct = 2859
	GarminProductVivoFit4                   GarminProduct = 2878
	GarminProductFr645                      GarminProduct = 2886
	GarminProductFr645m                     GarminProduct = 2888
	GarminProductFr30                       GarminProduct = 2891
	GarminProductFenix5sPlus                GarminProduct = 2900
	GarminProductEdge130                    GarminProduct = 2909
	GarminProductEdge1030Asia               GarminProduct = 2924
	GarminProductVivosmart4                 GarminProduct = 2927
	GarminProductVivoMoveHrAsia             GarminProduct = 2945
	GarminProductApproachX10                GarminProduct = 2962
	GarminProductFr30Asia                   GarminProduct = 2977
	GarminProductVivoactive3mW              GarminProduct = 2988
	GarminProductFr645Asia                  GarminProduct = 3003
	GarminProductFr645mAsia                 GarminProduct = 3004
	GarminProductEdgeExplore                GarminProduct = 3011
	GarminProductGpsmap66                   GarminProduct = 3028
	GarminProductApproachS10                GarminProduct = 3049
	GarminProductVivoactive3mL              GarminProduct = 3066
	GarminProductApproachG80                GarminProduct = 3085
	GarminProductEdge130Asia                GarminProduct = 3092
	GarminProductEdge1030Bontrager          GarminProduct = 3095
	GarminProductFenix5Plus                 GarminProduct = 3110
	GarminProductFenix5xPlus                GarminProduct = 3111
	GarminProductEdge520Plus                GarminProduct = 3112
	GarminProductFr945                      GarminProduct = 3113
	GarminProductEdge530                    GarminProduct = 3121
	GarminProductEdge830                    GarminProduct = 3122
	GarminProductInstinctEsports            GarminProduct = 3126
	GarminProductFenix5sPlusApac            GarminProduct = 3134
	GarminProductFenix5xPlusApac            GarminProduct = 3135
	GarminProductEdge520PlusApac            GarminProduct = 3142
	GarminProductDescentT1                  GarminProduct = 3143
	GarminProductFr235lAsia                 GarminProduct = 3144
	GarminProductFr245Asia                  GarminProduct = 3145
	GarminProductVivoActive3mApac           GarminProduct = 3163
	GarminProductGen3Bsm                    GarminProduct = 3192 // gen3 bike speed sensor
	GarminProductGen3Bcm                    GarminProduct = 3193 // gen3 bike cadence sensor
	GarminProductVivoSmart4Asia             GarminProduct = 3218
	GarminProductVivoactive4Small           GarminProduct = 3224
	GarminProductVivoactive4Large           GarminProduct = 3225
	GarminProductVenu                       GarminProduct = 3226
	GarminProductMarqDriver                 GarminProduct = 3246
	GarminProductMarqAviator                GarminProduct = 3247
	GarminProductMarqCaptain                GarminProduct = 3248
	GarminProductMarqCommander              GarminProduct = 3249
	GarminProductMarqExpedition             GarminProduct = 3250
	GarminProductMarqAthlete                GarminProduct = 3251
	GarminProductDescentMk2                 GarminProduct = 3258
	GarminProductGpsmap66i                  GarminProduct = 3284
	GarminProductFenix6SSport               GarminProduct = 3287
	GarminProductFenix6S                    GarminProduct = 3288
	GarminProductFenix6Sport                GarminProduct = 3289
	GarminProductFenix6                     GarminProduct = 3290
	GarminProductFenix6x                    GarminProduct = 3291
	GarminProductHrmDual                    GarminProduct = 3299 // HRM-Dual
	GarminProductHrmPro                     GarminProduct = 3300 // HRM-Pro
	GarminProductVivoMove3Premium           GarminProduct = 3308
	GarminProductApproachS40                GarminProduct = 3314
	GarminProductFr245mAsia                 GarminProduct = 3321
	GarminProductEdge530Apac                GarminProduct = 3349
	GarminProductEdge830Apac                GarminProduct = 3350
	GarminProductVivoMove3                  GarminProduct = 3378
	GarminProductVivoActive4SmallAsia       GarminProduct = 3387
	GarminProductVivoActive4LargeAsia       GarminProduct = 3388
	GarminProductVivoActive4OledAsia        GarminProduct = 3389
	GarminProductSwim2                      GarminProduct = 3405
	GarminProductMarqDriverAsia             GarminProduct = 3420
	GarminProductMarqAviatorAsia            GarminProduct = 3421
	GarminProductVivoMove3Asia              GarminProduct = 3422
	GarminProductFr945Asia                  GarminProduct = 3441
	GarminProductVivoActive3tChn            GarminProduct = 3446
	GarminProductMarqCaptainAsia            GarminProduct = 3448
	GarminProductMarqCommanderAsia          GarminProduct = 3449
	GarminProductMarqExpeditionAsia         GarminProduct = 3450
	GarminProductMarqAthleteAsia            GarminProduct = 3451
	GarminProductInstinctSolar              GarminProduct = 3466
	GarminProductFr45Asia                   GarminProduct = 3469
	GarminProductVivoactive3Daimler         GarminProduct = 3473
	GarminProductLegacyRey                  GarminProduct = 3498
	GarminProductLegacyDarthVader           GarminProduct = 3499
	GarminProductLegacyCaptainMarvel        GarminProduct = 3500
	GarminProductLegacyFirstAvenger         GarminProduct = 3501
	GarminProductFenix6sSportAsia           GarminProduct = 3512
	GarminProductFenix6sAsia                GarminProduct = 3513
	GarminProductFenix6SportAsia            GarminProduct = 3514
	GarminProductFenix6Asia                 GarminProduct = 3515
	GarminProductFenix6xAsia                GarminProduct = 3516
	GarminProductLegacyCaptainMarvelAsia    GarminProduct = 3535
	GarminProductLegacyFirstAvengerAsia     GarminProduct = 3536
	GarminProductLegacyReyAsia              GarminProduct = 3537
	GarminProductLegacyDarthVaderAsia       GarminProduct = 3538
	GarminProductDescentMk2s                GarminProduct = 3542
	GarminProductEdge130Plus                GarminProduct = 3558
	GarminProductEdge1030Plus               GarminProduct = 3570
	GarminProductRally200                   GarminProduct = 3578 // Rally 100/200 Power Meter Series
	GarminProductFr745                      GarminProduct = 3589
	GarminProductVenusq                     GarminProduct = 3600
	GarminProductLily                       GarminProduct = 3615
	GarminProductMarqAdventurer             GarminProduct = 3624
	GarminProductEnduro                     GarminProduct = 3638
	GarminProductSwim2Apac                  GarminProduct = 3639
	GarminProductMarqAdventurerAsia         GarminProduct = 3648
	GarminProductFr945Lte                   GarminProduct = 3652
	GarminProductDescentMk2Asia             GarminProduct = 3702 // Mk2 and Mk2i
	GarminProductVenu2                      GarminProduct = 3703
	GarminProductVenu2s                     GarminProduct = 3704
	GarminProductVenuDaimlerAsia            GarminProduct = 3737
	GarminProductMarqGolfer                 GarminProduct = 3739
	GarminProductVenuDaimler                GarminProduct = 3740
	GarminProductFr745Asia                  GarminProduct = 3794
	GarminProductLilyAsia                   GarminProduct = 3809
	GarminProductEdge1030PlusAsia           GarminProduct = 3812
	GarminProductEdge130PlusAsia            GarminProduct = 3813
	GarminProductApproachS12                GarminProduct = 3823
	GarminProductEnduroAsia                 GarminProduct = 3872
	GarminProductVenusqAsia                 GarminProduct = 3837
	GarminProductEdge1040                   GarminProduct = 3843
	GarminProductMarqGolferAsia             GarminProduct = 3850
	GarminProductVenu2Plus                  GarminProduct = 3851
	GarminProductGnss                       GarminProduct = 3865
	GarminProductFr55                       GarminProduct = 3869
	GarminProductInstinct2                  GarminProduct = 3888
	GarminProductFenix7s                    GarminProduct = 3905
	GarminProductFenix7                     GarminProduct = 3906
	GarminProductFenix7x                    GarminProduct = 3907
	GarminProductFenix7sApac                GarminProduct = 3908
	GarminProductFenix7Apac                 GarminProduct = 3909
	GarminProductFenix7xApac                GarminProduct = 3910
	GarminProductApproachG12                GarminProduct = 3927
	GarminProductDescentMk2sAsia            GarminProduct = 3930
	GarminProductApproachS42                GarminProduct = 3934
	GarminProductEpixGen2                   GarminProduct = 3943
	GarminProductEpixGen2Apac               GarminProduct = 3944
	GarminProductVenu2sAsia                 GarminProduct = 3949
	GarminProductVenu2Asia                  GarminProduct = 3950
	GarminProductFr945LteAsia               GarminProduct = 3978
	GarminProductVivoMoveSport              GarminProduct = 3982
	GarminProductVivomoveTrend              GarminProduct = 3983
	GarminProductApproachS12Asia            GarminProduct = 3986
	GarminProductFr255Music                 GarminProduct = 3990
	GarminProductFr255SmallMusic            GarminProduct = 3991
	GarminProductFr255                      GarminProduct = 3992
	GarminProductFr255Small                 GarminProduct = 3993
	GarminProductApproachG12Asia            GarminProduct = 4001
	GarminProductApproachS42Asia            GarminProduct = 4002
	GarminProductDescentG1                  GarminProduct = 4005
	GarminProductVenu2PlusAsia              GarminProduct = 4017
	GarminProductFr955                      GarminProduct = 4024
	GarminProductFr55Asia                   GarminProduct = 4033
	GarminProductEdge540                    GarminProduct = 4061
	GarminProductEdge840                    GarminProduct = 4062
	GarminProductVivosmart5                 GarminProduct = 4063
	GarminProductInstinct2Asia              GarminProduct = 4071
	GarminProductMarqGen2                   GarminProduct = 4105 // Adventurer, Athlete, Captain, Golfer
	GarminProductVenusq2                    GarminProduct = 4115
	GarminProductVenusq2music               GarminProduct = 4116
	GarminProductMarqGen2Aviator            GarminProduct = 4124
	GarminProductD2AirX10                   GarminProduct = 4125
	GarminProductHrmProPlus                 GarminProduct = 4130
	GarminProductDescentG1Asia              GarminProduct = 4132
	GarminProductTactix7                    GarminProduct = 4135
	GarminProductInstinctCrossover          GarminProduct = 4155
	GarminProductEdgeExplore2               GarminProduct = 4169
	GarminProductApproachS70                GarminProduct = 4233
	GarminProductFr265Large                 GarminProduct = 4257
	GarminProductFr265Small                 GarminProduct = 4258
	GarminProductTacxNeoSmart               GarminProduct = 4265 // Neo Smart, Tacx
	GarminProductTacxNeo2Smart              GarminProduct = 4266 // Neo 2 Smart, Tacx
	GarminProductTacxNeo2TSmart             GarminProduct = 4267 // Neo 2T Smart, Tacx
	GarminProductTacxNeoSmartBike           GarminProduct = 4268 // Neo Smart Bike, Tacx
	GarminProductTacxSatoriSmart            GarminProduct = 4269 // Satori Smart, Tacx
	GarminProductTacxFlowSmart              GarminProduct = 4270 // Flow Smart, Tacx
	GarminProductTacxVortexSmart            GarminProduct = 4271 // Vortex Smart, Tacx
	GarminProductTacxBushidoSmart           GarminProduct = 4272 // Bushido Smart, Tacx
	GarminProductTacxGeniusSmart            GarminProduct = 4273 // Genius Smart, Tacx
	GarminProductTacxFluxFluxSSmart         GarminProduct = 4274 // Flux/Flux S Smart, Tacx
	GarminProductTacxFlux2Smart             GarminProduct = 4275 // Flux 2 Smart, Tacx
	GarminProductTacxMagnum                 GarminProduct = 4276 // Magnum, Tacx
	GarminProductEdge1040Asia               GarminProduct = 4305
	GarminProductEpixGen2Pro42              GarminProduct = 4312
	GarminProductEpixGen2Pro47              GarminProduct = 4313
	GarminProductEpixGen2Pro51              GarminProduct = 4314
	GarminProductFr965                      GarminProduct = 4315
	GarminProductEnduro2                    GarminProduct = 4341
	GarminProductFenix7ProSolar             GarminProduct = 4375
	GarminProductInstinct2x                 GarminProduct = 4394
	GarminProductDescentT2                  GarminProduct = 4442
	GarminProductSdm4                       GarminProduct = 10007 // SDM4 footpod
	GarminProductEdgeRemote                 GarminProduct = 10014
	GarminProductTacxTrainingAppWin         GarminProduct = 20533
	GarminProductTacxTrainingAppMac         GarminProduct = 20534
	GarminProductTacxTrainingAppMacCatalyst GarminProduct = 20565
	GarminProductTrainingCenter             GarminProduct = 20119
	GarminProductTacxTrainingAppAndroid     GarminProduct = 30045
	GarminProductTacxTrainingAppIos         GarminProduct = 30046
	GarminProductTacxTrainingAppLegacy      GarminProduct = 30047
	GarminProductConnectiqSimulator         GarminProduct = 65531
	GarminProductAndroidAntplusPlugin       GarminProduct = 65532
	GarminProductConnect                    GarminProduct = 65534 // Garmin Connect website
	GarminProductInvalid                    GarminProduct = 0xFFFF
)

// GasConsumptionRateType represents the gas_consumption_rate_type FIT type.
type GasConsumptionRateType byte

const (
	GasConsumptionRateTypePressureSac GasConsumptionRateType = 0 // Pressure-based Surface Air Consumption
	GasConsumptionRateTypeVolumeSac   GasConsumptionRateType = 1 // Volumetric Surface Air Consumption
	GasConsumptionRateTypeRmv         GasConsumptionRateType = 2 // Respiratory Minute Volume
	GasConsumptionRateTypeInvalid     GasConsumptionRateType = 0xFF
)

// Gender represents the gender FIT type.
type Gender byte

const (
	GenderFemale  Gender = 0
	GenderMale    Gender = 1
	GenderInvalid Gender = 0xFF
)

// Goal represents the goal FIT type.
type Goal byte

const (
	GoalTime          Goal = 0
	GoalDistance      Goal = 1
	GoalCalories      Goal = 2
	GoalFrequency     Goal = 3
	GoalSteps         Goal = 4
	GoalAscent        Goal = 5
	GoalActiveMinutes Goal = 6
	GoalInvalid       Goal = 0xFF
)

// GoalRecurrence represents the goal_recurrence FIT type.
type GoalRecurrence byte

const (
	GoalRecurrenceOff     GoalRecurrence = 0
	GoalRecurrenceDaily   GoalRecurrence = 1
	GoalRecurrenceWeekly  GoalRecurrence = 2
	GoalRecurrenceMonthly GoalRecurrence = 3
	GoalRecurrenceYearly  GoalRecurrence = 4
	GoalRecurrenceCustom  GoalRecurrence = 5
	GoalRecurrenceInvalid GoalRecurrence = 0xFF
)

// GoalSource represents the goal_source FIT type.
type GoalSource byte

const (
	GoalSourceAuto      GoalSource = 0 // Device generated
	GoalSourceCommunity GoalSource = 1 // Social network sourced goal
	GoalSourceUser      GoalSource = 2 // Manually generated
	GoalSourceInvalid   GoalSource = 0xFF
)

// HipRaiseExerciseName represents the hip_raise_exercise_name FIT type.
type HipRaiseExerciseName uint16

const (
	HipRaiseExerciseNameBarbellHipThrustOnFloor                         HipRaiseExerciseName = 0
	HipRaiseExerciseNameBarbellHipThrustWithBench                       HipRaiseExerciseName = 1
	HipRaiseExerciseNameBentKneeSwissBallReverseHipRaise                HipRaiseExerciseName = 2
	HipRaiseExerciseNameWeightedBentKneeSwissBallReverseHipRaise        HipRaiseExerciseName = 3
	HipRaiseExerciseNameBridgeWithLegExtension                          HipRaiseExerciseName = 4
	HipRaiseExerciseNameWeightedBridgeWithLegExtension                  HipRaiseExerciseName = 5
	HipRaiseExerciseNameClamBridge                                      HipRaiseExerciseName = 6
	HipRaiseExerciseNameFrontKickTabletop                               HipRaiseExerciseName = 7
	HipRaiseExerciseNameWeightedFrontKickTabletop                       HipRaiseExerciseName = 8
	HipRaiseExerciseNameHipExtensionAndCross                            HipRaiseExerciseName = 9
	HipRaiseExerciseNameWeightedHipExtensionAndCross                    HipRaiseExerciseName = 10
	HipRaiseExerciseNameHipRaise                                        HipRaiseExerciseName = 11
	HipRaiseExerciseNameWeightedHipRaise                                HipRaiseExerciseName = 12
	HipRaiseExerciseNameHipRaiseWithFeetOnSwissBall                     HipRaiseExerciseName = 13
	HipRaiseExerciseNameWeightedHipRaiseWithFeetOnSwissBall             HipRaiseExerciseName = 14
	HipRaiseExerciseNameHipRaiseWithHeadOnBosuBall                      HipRaiseExerciseName = 15
	HipRaiseExerciseNameWeightedHipRaiseWithHeadOnBosuBall              HipRaiseExerciseName = 16
	HipRaiseExerciseNameHipRaiseWithHeadOnSwissBall                     HipRaiseExerciseName = 17
	HipRaiseExerciseNameWeightedHipRaiseWithHeadOnSwissBall             HipRaiseExerciseName = 18
	HipRaiseExerciseNameHipRaiseWithKneeSqueeze                         HipRaiseExerciseName = 19
	HipRaiseExerciseNameWeightedHipRaiseWithKneeSqueeze                 HipRaiseExerciseName = 20
	HipRaiseExerciseNameInclineRearLegExtension                         HipRaiseExerciseName = 21
	HipRaiseExerciseNameWeightedInclineRearLegExtension                 HipRaiseExerciseName = 22
	HipRaiseExerciseNameKettlebellSwing                                 HipRaiseExerciseName = 23
	HipRaiseExerciseNameMarchingHipRaise                                HipRaiseExerciseName = 24
	HipRaiseExerciseNameWeightedMarchingHipRaise                        HipRaiseExerciseName = 25
	HipRaiseExerciseNameMarchingHipRaiseWithFeetOnASwissBall            HipRaiseExerciseName = 26
	HipRaiseExerciseNameWeightedMarchingHipRaiseWithFeetOnASwissBall    HipRaiseExerciseName = 27
	HipRaiseExerciseNameReverseHipRaise                                 HipRaiseExerciseName = 28
	HipRaiseExerciseNameWeightedReverseHipRaise                         HipRaiseExerciseName = 29
	HipRaiseExerciseNameSingleLegHipRaise                               HipRaiseExerciseName = 30
	HipRaiseExerciseNameWeightedSingleLegHipRaise                       HipRaiseExerciseName = 31
	HipRaiseExerciseNameSingleLegHipRaiseWithFootOnBench                HipRaiseExerciseName = 32
	HipRaiseExerciseNameWeightedSingleLegHipRaiseWithFootOnBench        HipRaiseExerciseName = 33
	HipRaiseExerciseNameSingleLegHipRaiseWithFootOnBosuBall             HipRaiseExerciseName = 34
	HipRaiseExerciseNameWeightedSingleLegHipRaiseWithFootOnBosuBall     HipRaiseExerciseName = 35
	HipRaiseExerciseNameSingleLegHipRaiseWithFootOnFoamRoller           HipRaiseExerciseName = 36
	HipRaiseExerciseNameWeightedSingleLegHipRaiseWithFootOnFoamRoller   HipRaiseExerciseName = 37
	HipRaiseExerciseNameSingleLegHipRaiseWithFootOnMedicineBall         HipRaiseExerciseName = 38
	HipRaiseExerciseNameWeightedSingleLegHipRaiseWithFootOnMedicineBall HipRaiseExerciseName = 39
	HipRaiseExerciseNameSingleLegHipRaiseWithHeadOnBosuBall             HipRaiseExerciseName = 40
	HipRaiseExerciseNameWeightedSingleLegHipRaiseWithHeadOnBosuBall     HipRaiseExerciseName = 41
	HipRaiseExerciseNameWeightedClamBridge                              HipRaiseExerciseName = 42
	HipRaiseExerciseNameSingleLegSwissBallHipRaiseAndLegCurl            HipRaiseExerciseName = 43
	HipRaiseExerciseNameClams                                           HipRaiseExerciseName = 44
	HipRaiseExerciseNameInnerThighCircles                               HipRaiseExerciseName = 45 // Deprecated do not use
	HipRaiseExerciseNameInnerThighSideLift                              HipRaiseExerciseName = 46 // Deprecated do not use
	HipRaiseExerciseNameLegCircles                                      HipRaiseExerciseName = 47
	HipRaiseExerciseNameLegLift                                         HipRaiseExerciseName = 48
	HipRaiseExerciseNameLegLiftInExternalRotation                       HipRaiseExerciseName = 49
	HipRaiseExerciseNameInvalid                                         HipRaiseExerciseName = 0xFFFF
)

// HipStabilityExerciseName represents the hip_stability_exercise_name FIT type.
type HipStabilityExerciseName uint16

const (
	HipStabilityExerciseNameBandSideLyingLegRaise             HipStabilityExerciseName = 0
	HipStabilityExerciseNameDeadBug                           HipStabilityExerciseName = 1
	HipStabilityExerciseNameWeightedDeadBug                   HipStabilityExerciseName = 2
	HipStabilityExerciseNameExternalHipRaise                  HipStabilityExerciseName = 3
	HipStabilityExerciseNameWeightedExternalHipRaise          HipStabilityExerciseName = 4
	HipStabilityExerciseNameFireHydrantKicks                  HipStabilityExerciseName = 5
	HipStabilityExerciseNameWeightedFireHydrantKicks          HipStabilityExerciseName = 6
	HipStabilityExerciseNameHipCircles                        HipStabilityExerciseName = 7
	HipStabilityExerciseNameWeightedHipCircles                HipStabilityExerciseName = 8
	HipStabilityExerciseNameInnerThighLift                    HipStabilityExerciseName = 9
	HipStabilityExerciseNameWeightedInnerThighLift            HipStabilityExerciseName = 10
	HipStabilityExerciseNameLateralWalksWithBandAtAnkles      HipStabilityExerciseName = 11
	HipStabilityExerciseNamePretzelSideKick                   HipStabilityExerciseName = 12
	HipStabilityExerciseNameWeightedPretzelSideKick           HipStabilityExerciseName = 13
	HipStabilityExerciseNameProneHipInternalRotation          HipStabilityExerciseName = 14
	HipStabilityExerciseNameWeightedProneHipInternalRotation  HipStabilityExerciseName = 15
	HipStabilityExerciseNameQuadruped                         HipStabilityExerciseName = 16
	HipStabilityExerciseNameQuadrupedHipExtension             HipStabilityExerciseName = 17
	HipStabilityExerciseNameWeightedQuadrupedHipExtension     HipStabilityExerciseName = 18
	HipStabilityExerciseNameQuadrupedWithLegLift              HipStabilityExerciseName = 19
	HipStabilityExerciseNameWeightedQuadrupedWithLegLift      HipStabilityExerciseName = 20
	HipStabilityExerciseNameSideLyingLegRaise                 HipStabilityExerciseName = 21
	HipStabilityExerciseNameWeightedSideLyingLegRaise         HipStabilityExerciseName = 22
	HipStabilityExerciseNameSlidingHipAdduction               HipStabilityExerciseName = 23
	HipStabilityExerciseNameWeightedSlidingHipAdduction       HipStabilityExerciseName = 24
	HipStabilityExerciseNameStandingAdduction                 HipStabilityExerciseName = 25
	HipStabilityExerciseNameWeightedStandingAdduction         HipStabilityExerciseName = 26
	HipStabilityExerciseNameStandingCableHipAbduction         HipStabilityExerciseName = 27
	HipStabilityExerciseNameStandingHipAbduction              HipStabilityExerciseName = 28
	HipStabilityExerciseNameWeightedStandingHipAbduction      HipStabilityExerciseName = 29
	HipStabilityExerciseNameStandingRearLegRaise              HipStabilityExerciseName = 30
	HipStabilityExerciseNameWeightedStandingRearLegRaise      HipStabilityExerciseName = 31
	HipStabilityExerciseNameSupineHipInternalRotation         HipStabilityExerciseName = 32
	HipStabilityExerciseNameWeightedSupineHipInternalRotation HipStabilityExerciseName = 33
	HipStabilityExerciseNameInvalid                           HipStabilityExerciseName = 0xFFFF
)

// HipSwingExerciseName represents the hip_swing_exercise_name FIT type.
type HipSwingExerciseName uint16

const (
	HipSwingExerciseNameSingleArmKettlebellSwing HipSwingExerciseName = 0
	HipSwingExerciseNameSingleArmDumbbellSwing   HipSwingExerciseName = 1
	HipSwingExerciseNameStepOutSwing             HipSwingExerciseName = 2
	HipSwingExerciseNameInvalid                  HipSwingExerciseName = 0xFFFF
)

// HrType represents the hr_type FIT type.
type HrType byte

const (
	HrTypeNormal    HrType = 0
	HrTypeIrregular HrType = 1
	HrTypeInvalid   HrType = 0xFF
)

// HrZoneCalc represents the hr_zone_calc FIT type.
type HrZoneCalc byte

const (
	HrZoneCalcCustom       HrZoneCalc = 0
	HrZoneCalcPercentMaxHr HrZoneCalc = 1
	HrZoneCalcPercentHrr   HrZoneCalc = 2
	HrZoneCalcPercentLthr  HrZoneCalc = 3
	HrZoneCalcInvalid      HrZoneCalc = 0xFF
)

// HrvStatus represents the hrv_status FIT type.
type HrvStatus byte

const (
	HrvStatusNone       HrvStatus = 0
	HrvStatusPoor       HrvStatus = 1
	HrvStatusLow        HrvStatus = 2
	HrvStatusUnbalanced HrvStatus = 3
	HrvStatusBalanced   HrvStatus = 4
	HrvStatusInvalid    HrvStatus = 0xFF
)

// HyperextensionExerciseName represents the hyperextension_exercise_name FIT type.
type HyperextensionExerciseName uint16

const (
	HyperextensionExerciseNameBackExtensionWithOppositeArmAndLegReach         HyperextensionExerciseName = 0
	HyperextensionExerciseNameWeightedBackExtensionWithOppositeArmAndLegReach HyperextensionExerciseName = 1
	HyperextensionExerciseNameBaseRotations                                   HyperextensionExerciseName = 2
	HyperextensionExerciseNameWeightedBaseRotations                           HyperextensionExerciseName = 3
	HyperextensionExerciseNameBentKneeReverseHyperextension                   HyperextensionExerciseName = 4
	HyperextensionExerciseNameWeightedBentKneeReverseHyperextension           HyperextensionExerciseName = 5
	HyperextensionExerciseNameHollowHoldAndRoll                               HyperextensionExerciseName = 6
	HyperextensionExerciseNameWeightedHollowHoldAndRoll                       HyperextensionExerciseName = 7
	HyperextensionExerciseNameKicks                                           HyperextensionExerciseName = 8
	HyperextensionExerciseNameWeightedKicks                                   HyperextensionExerciseName = 9
	HyperextensionExerciseNameKneeRaises                                      HyperextensionExerciseName = 10
	HyperextensionExerciseNameWeightedKneeRaises                              HyperextensionExerciseName = 11
	HyperextensionExerciseNameKneelingSuperman                                HyperextensionExerciseName = 12
	HyperextensionExerciseNameWeightedKneelingSuperman                        HyperextensionExerciseName = 13
	HyperextensionExerciseNameLatPullDownWithRow                              HyperextensionExerciseName = 14
	HyperextensionExerciseNameMedicineBallDeadliftToReach                     HyperextensionExerciseName = 15
	HyperextensionExerciseNameOneArmOneLegRow                                 HyperextensionExerciseName = 16
	HyperextensionExerciseNameOneArmRowWithBand                               HyperextensionExerciseName = 17
	HyperextensionExerciseNameOverheadLungeWithMedicineBall                   HyperextensionExerciseName = 18
	HyperextensionExerciseNamePlankKneeTucks                                  HyperextensionExerciseName = 19
	HyperextensionExerciseNameWeightedPlankKneeTucks                          HyperextensionExerciseName = 20
	HyperextensionExerciseNameSideStep                                        HyperextensionExerciseName = 21
	HyperextensionExerciseNameWeightedSideStep                                HyperextensionExerciseName = 22
	HyperextensionExerciseNameSingleLegBackExtension                          HyperextensionExerciseName = 23
	HyperextensionExerciseNameWeightedSingleLegBackExtension                  HyperextensionExerciseName = 24
	HyperextensionExerciseNameSpineExtension                                  HyperextensionExerciseName = 25
	HyperextensionExerciseNameWeightedSpineExtension                          HyperextensionExerciseName = 26
	HyperextensionExerciseNameStaticBackExtension                             HyperextensionExerciseName = 27
	HyperextensionExerciseNameWeightedStaticBackExtension                     HyperextensionExerciseName = 28
	HyperextensionExerciseNameSupermanFromFloor                               HyperextensionExerciseName = 29
	HyperextensionExerciseNameWeightedSupermanFromFloor                       HyperextensionExerciseName = 30
	HyperextensionExerciseNameSwissBallBackExtension                          HyperextensionExerciseName = 31
	HyperextensionExerciseNameWeightedSwissBallBackExtension                  HyperextensionExerciseName = 32
	HyperextensionExerciseNameSwissBallHyperextension                         HyperextensionExerciseName = 33
	HyperextensionExerciseNameWeightedSwissBallHyperextension                 HyperextensionExerciseName = 34
	HyperextensionExerciseNameSwissBallOppositeArmAndLegLift                  HyperextensionExerciseName = 35
	HyperextensionExerciseNameWeightedSwissBallOppositeArmAndLegLift          HyperextensionExerciseName = 36
	HyperextensionExerciseNameSupermanOnSwissBall                             HyperextensionExerciseName = 37
	HyperextensionExerciseNameCobra                                           HyperextensionExerciseName = 38
	HyperextensionExerciseNameSupineFloorBarre                                HyperextensionExerciseName = 39 // Deprecated do not use
	HyperextensionExerciseNameInvalid                                         HyperextensionExerciseName = 0xFFFF
)

// Intensity represents the intensity FIT type.
type Intensity byte

const (
	IntensityActive   Intensity = 0
	IntensityRest     Intensity = 1
	IntensityWarmup   Intensity = 2
	IntensityCooldown Intensity = 3
	IntensityRecovery Intensity = 4
	IntensityInterval Intensity = 5
	IntensityOther    Intensity = 6
	IntensityInvalid  Intensity = 0xFF
)

// Language represents the language FIT type.
type Language byte

const (
	LanguageEnglish             Language = 0
	LanguageFrench              Language = 1
	LanguageItalian             Language = 2
	LanguageGerman              Language = 3
	LanguageSpanish             Language = 4
	LanguageCroatian            Language = 5
	LanguageCzech               Language = 6
	LanguageDanish              Language = 7
	LanguageDutch               Language = 8
	LanguageFinnish             Language = 9
	LanguageGreek               Language = 10
	LanguageHungarian           Language = 11
	LanguageNorwegian           Language = 12
	LanguagePolish              Language = 13
	LanguagePortuguese          Language = 14
	LanguageSlovakian           Language = 15
	LanguageSlovenian           Language = 16
	LanguageSwedish             Language = 17
	LanguageRussian             Language = 18
	LanguageTurkish             Language = 19
	LanguageLatvian             Language = 20
	LanguageUkrainian           Language = 21
	LanguageArabic              Language = 22
	LanguageFarsi               Language = 23
	LanguageBulgarian           Language = 24
	LanguageRomanian            Language = 25
	LanguageChinese             Language = 26
	LanguageJapanese            Language = 27
	LanguageKorean              Language = 28
	LanguageTaiwanese           Language = 29
	LanguageThai                Language = 30
	LanguageHebrew              Language = 31
	LanguageBrazilianPortuguese Language = 32
	LanguageIndonesian          Language = 33
	LanguageMalaysian           Language = 34
	LanguageVietnamese          Language = 35
	LanguageBurmese             Language = 36
	LanguageMongolian           Language = 37
	LanguageCustom              Language = 254
	LanguageInvalid             Language = 0xFF
)

// LanguageBits0 represents the language_bits_0 FIT type.
type LanguageBits0 uint8

const (
	LanguageBits0English  LanguageBits0 = 0x01
	LanguageBits0French   LanguageBits0 = 0x02
	LanguageBits0Italian  LanguageBits0 = 0x04
	LanguageBits0German   LanguageBits0 = 0x08
	LanguageBits0Spanish  LanguageBits0 = 0x10
	LanguageBits0Croatian LanguageBits0 = 0x20
	LanguageBits0Czech    LanguageBits0 = 0x40
	LanguageBits0Danish   LanguageBits0 = 0x80
	LanguageBits0Invalid  LanguageBits0 = 0x00
)

// LanguageBits1 represents the language_bits_1 FIT type.
type LanguageBits1 uint8

const (
	LanguageBits1Dutch      LanguageBits1 = 0x01
	LanguageBits1Finnish    LanguageBits1 = 0x02
	LanguageBits1Greek      LanguageBits1 = 0x04
	LanguageBits1Hungarian  LanguageBits1 = 0x08
	LanguageBits1Norwegian  LanguageBits1 = 0x10
	LanguageBits1Polish     LanguageBits1 = 0x20
	LanguageBits1Portuguese LanguageBits1 = 0x40
	LanguageBits1Slovakian  LanguageBits1 = 0x80
	LanguageBits1Invalid    LanguageBits1 = 0x00
)

// LanguageBits2 represents the language_bits_2 FIT type.
type LanguageBits2 uint8

const (
	LanguageBits2Slovenian LanguageBits2 = 0x01
	LanguageBits2Swedish   LanguageBits2 = 0x02
	LanguageBits2Russian   LanguageBits2 = 0x04
	LanguageBits2Turkish   LanguageBits2 = 0x08
	LanguageBits2Latvian   LanguageBits2 = 0x10
	LanguageBits2Ukrainian LanguageBits2 = 0x20
	LanguageBits2Arabic    LanguageBits2 = 0x40
	LanguageBits2Farsi     LanguageBits2 = 0x80
	LanguageBits2Invalid   LanguageBits2 = 0x00
)

// LanguageBits3 represents the language_bits_3 FIT type.
type LanguageBits3 uint8

const (
	LanguageBits3Bulgarian LanguageBits3 = 0x01
	LanguageBits3Romanian  LanguageBits3 = 0x02
	LanguageBits3Chinese   LanguageBits3 = 0x04
	LanguageBits3Japanese  LanguageBits3 = 0x08
	LanguageBits3Korean    LanguageBits3 = 0x10
	LanguageBits3Taiwanese LanguageBits3 = 0x20
	LanguageBits3Thai      LanguageBits3 = 0x40
	LanguageBits3Hebrew    LanguageBits3 = 0x80
	LanguageBits3Invalid   LanguageBits3 = 0x00
)

// LanguageBits4 represents the language_bits_4 FIT type.
type LanguageBits4 uint8

const (
	LanguageBits4BrazilianPortuguese LanguageBits4 = 0x01
	LanguageBits4Indonesian          LanguageBits4 = 0x02
	LanguageBits4Malaysian           LanguageBits4 = 0x04
	LanguageBits4Vietnamese          LanguageBits4 = 0x08
	LanguageBits4Burmese             LanguageBits4 = 0x10
	LanguageBits4Mongolian           LanguageBits4 = 0x20
	LanguageBits4Invalid             LanguageBits4 = 0x00
)

// LapTrigger represents the lap_trigger FIT type.
type LapTrigger byte

const (
	LapTriggerManual           LapTrigger = 0
	LapTriggerTime             LapTrigger = 1
	LapTriggerDistance         LapTrigger = 2
	LapTriggerPositionStart    LapTrigger = 3
	LapTriggerPositionLap      LapTrigger = 4
	LapTriggerPositionWaypoint LapTrigger = 5
	LapTriggerPositionMarked   LapTrigger = 6
	LapTriggerSessionEnd       LapTrigger = 7
	LapTriggerFitnessEquipment LapTrigger = 8
	LapTriggerInvalid          LapTrigger = 0xFF
)

// LateralRaiseExerciseName represents the lateral_raise_exercise_name FIT type.
type LateralRaiseExerciseName uint16

const (
	LateralRaiseExerciseName45DegreeCableExternalRotation         LateralRaiseExerciseName = 0
	LateralRaiseExerciseNameAlternatingLateralRaiseWithStaticHold LateralRaiseExerciseName = 1
	LateralRaiseExerciseNameBarMuscleUp                           LateralRaiseExerciseName = 2
	LateralRaiseExerciseNameBentOverLateralRaise                  LateralRaiseExerciseName = 3
	LateralRaiseExerciseNameCableDiagonalRaise                    LateralRaiseExerciseName = 4
	LateralRaiseExerciseNameCableFrontRaise                       LateralRaiseExerciseName = 5
	LateralRaiseExerciseNameCalorieRow                            LateralRaiseExerciseName = 6
	LateralRaiseExerciseNameComboShoulderRaise                    LateralRaiseExerciseName = 7
	LateralRaiseExerciseNameDumbbellDiagonalRaise                 LateralRaiseExerciseName = 8
	LateralRaiseExerciseNameDumbbellVRaise                        LateralRaiseExerciseName = 9
	LateralRaiseExerciseNameFrontRaise                            LateralRaiseExerciseName = 10
	LateralRaiseExerciseNameLeaningDumbbellLateralRaise           LateralRaiseExerciseName = 11
	LateralRaiseExerciseNameLyingDumbbellRaise                    LateralRaiseExerciseName = 12
	LateralRaiseExerciseNameMuscleUp                              LateralRaiseExerciseName = 13
	LateralRaiseExerciseNameOneArmCableLateralRaise               LateralRaiseExerciseName = 14
	LateralRaiseExerciseNameOverhandGripRearLateralRaise          LateralRaiseExerciseName = 15
	LateralRaiseExerciseNamePlateRaises                           LateralRaiseExerciseName = 16
	LateralRaiseExerciseNameRingDip                               LateralRaiseExerciseName = 17
	LateralRaiseExerciseNameWeightedRingDip                       LateralRaiseExerciseName = 18
	LateralRaiseExerciseNameRingMuscleUp                          LateralRaiseExerciseName = 19
	LateralRaiseExerciseNameWeightedRingMuscleUp                  LateralRaiseExerciseName = 20
	LateralRaiseExerciseNameRopeClimb                             LateralRaiseExerciseName = 21
	LateralRaiseExerciseNameWeightedRopeClimb                     LateralRaiseExerciseName = 22
	LateralRaiseExerciseNameScaption                              LateralRaiseExerciseName = 23
	LateralRaiseExerciseNameSeatedLateralRaise                    LateralRaiseExerciseName = 24
	LateralRaiseExerciseNameSeatedRearLateralRaise                LateralRaiseExerciseName = 25
	LateralRaiseExerciseNameSideLyingLateralRaise                 LateralRaiseExerciseName = 26
	LateralRaiseExerciseNameStandingLift                          LateralRaiseExerciseName = 27
	LateralRaiseExerciseNameSuspendedRow                          LateralRaiseExerciseName = 28
	LateralRaiseExerciseNameUnderhandGripRearLateralRaise         LateralRaiseExerciseName = 29
	LateralRaiseExerciseNameWallSlide                             LateralRaiseExerciseName = 30
	LateralRaiseExerciseNameWeightedWallSlide                     LateralRaiseExerciseName = 31
	LateralRaiseExerciseNameArmCircles                            LateralRaiseExerciseName = 32
	LateralRaiseExerciseNameShavingTheHead                        LateralRaiseExerciseName = 33
	LateralRaiseExerciseNameInvalid                               LateralRaiseExerciseName = 0xFFFF
)

// LeftRightBalance represents the left_right_balance FIT type.
type LeftRightBalance uint8

const (
	LeftRightBalanceMask    LeftRightBalance = 0x7F // % contribution
	LeftRightBalanceRight   LeftRightBalance = 0x80 // data corresponds to right if set, otherwise unknown
	LeftRightBalanceInvalid LeftRightBalance = 0xFF
)

// LeftRightBalance100 represents the left_right_balance_100 FIT type.
type LeftRightBalance100 uint16

const (
	LeftRightBalance100Mask    LeftRightBalance100 = 0x3FFF // % contribution scaled by 100
	LeftRightBalance100Right   LeftRightBalance100 = 0x8000 // data corresponds to right if set, otherwise unknown
	LeftRightBalance100Invalid LeftRightBalance100 = 0xFFFF
)

// LegCurlExerciseName represents the leg_curl_exercise_name FIT type.
type LegCurlExerciseName uint16

const (
	LegCurlExerciseNameLegCurl                     LegCurlExerciseName = 0
	LegCurlExerciseNameWeightedLegCurl             LegCurlExerciseName = 1
	LegCurlExerciseNameGoodMorning                 LegCurlExerciseName = 2
	LegCurlExerciseNameSeatedBarbellGoodMorning    LegCurlExerciseName = 3
	LegCurlExerciseNameSingleLegBarbellGoodMorning LegCurlExerciseName = 4
	LegCurlExerciseNameSingleLegSlidingLegCurl     LegCurlExerciseName = 5
	LegCurlExerciseNameSlidingLegCurl              LegCurlExerciseName = 6
	LegCurlExerciseNameSplitBarbellGoodMorning     LegCurlExerciseName = 7
	LegCurlExerciseNameSplitStanceExtension        LegCurlExerciseName = 8
	LegCurlExerciseNameStaggeredStanceGoodMorning  LegCurlExerciseName = 9
	LegCurlExerciseNameSwissBallHipRaiseAndLegCurl LegCurlExerciseName = 10
	LegCurlExerciseNameZercherGoodMorning          LegCurlExerciseName = 11
	LegCurlExerciseNameInvalid                     LegCurlExerciseName = 0xFFFF
)

// LegRaiseExerciseName represents the leg_raise_exercise_name FIT type.
type LegRaiseExerciseName uint16

const (
	LegRaiseExerciseNameHangingKneeRaise                   LegRaiseExerciseName = 0
	LegRaiseExerciseNameHangingLegRaise                    LegRaiseExerciseName = 1
	LegRaiseExerciseNameWeightedHangingLegRaise            LegRaiseExerciseName = 2
	LegRaiseExerciseNameHangingSingleLegRaise              LegRaiseExerciseName = 3
	LegRaiseExerciseNameWeightedHangingSingleLegRaise      LegRaiseExerciseName = 4
	LegRaiseExerciseNameKettlebellLegRaises                LegRaiseExerciseName = 5
	LegRaiseExerciseNameLegLoweringDrill                   LegRaiseExerciseName = 6
	LegRaiseExerciseNameWeightedLegLoweringDrill           LegRaiseExerciseName = 7
	LegRaiseExerciseNameLyingStraightLegRaise              LegRaiseExerciseName = 8
	LegRaiseExerciseNameWeightedLyingStraightLegRaise      LegRaiseExerciseName = 9
	LegRaiseExerciseNameMedicineBallLegDrops               LegRaiseExerciseName = 10
	LegRaiseExerciseNameQuadrupedLegRaise                  LegRaiseExerciseName = 11
	LegRaiseExerciseNameWeightedQuadrupedLegRaise          LegRaiseExerciseName = 12
	LegRaiseExerciseNameReverseLegRaise                    LegRaiseExerciseName = 13
	LegRaiseExerciseNameWeightedReverseLegRaise            LegRaiseExerciseName = 14
	LegRaiseExerciseNameReverseLegRaiseOnSwissBall         LegRaiseExerciseName = 15
	LegRaiseExerciseNameWeightedReverseLegRaiseOnSwissBall LegRaiseExerciseName = 16
	LegRaiseExerciseNameSingleLegLoweringDrill             LegRaiseExerciseName = 17
	LegRaiseExerciseNameWeightedSingleLegLoweringDrill     LegRaiseExerciseName = 18
	LegRaiseExerciseNameWeightedHangingKneeRaise           LegRaiseExerciseName = 19
	LegRaiseExerciseNameLateralStepover                    LegRaiseExerciseName = 20
	LegRaiseExerciseNameWeightedLateralStepover            LegRaiseExerciseName = 21
	LegRaiseExerciseNameInvalid                            LegRaiseExerciseName = 0xFFFF
)

// LengthType represents the length_type FIT type.
type LengthType byte

const (
	LengthTypeIdle    LengthType = 0 // Rest period. Length with no strokes
	LengthTypeActive  LengthType = 1 // Length with strokes.
	LengthTypeInvalid LengthType = 0xFF
)

// LocalDeviceType represents the local_device_type FIT type.
type LocalDeviceType uint8

const (
	LocalDeviceTypeGps           LocalDeviceType = 0  // Onboard gps receiver
	LocalDeviceTypeGlonass       LocalDeviceType = 1  // Onboard glonass receiver
	LocalDeviceTypeGpsGlonass    LocalDeviceType = 2  // Onboard gps glonass receiver
	LocalDeviceTypeAccelerometer LocalDeviceType = 3  // Onboard sensor
	LocalDeviceTypeBarometer     LocalDeviceType = 4  // Onboard sensor
	LocalDeviceTypeTemperature   LocalDeviceType = 5  // Onboard sensor
	LocalDeviceTypeWhr           LocalDeviceType = 10 // Onboard wrist HR sensor
	LocalDeviceTypeSensorHub     LocalDeviceType = 12 // Onboard software package
	LocalDeviceTypeInvalid       LocalDeviceType = 0xFF
)

// LocaltimeIntoDay represents the localtime_into_day FIT type.
type LocaltimeIntoDay uint32

const (
	LocaltimeIntoDayInvalid LocaltimeIntoDay = 0xFFFFFFFF
)

// LungeExerciseName represents the lunge_exercise_name FIT type.
type LungeExerciseName uint16

const (
	LungeExerciseNameOverheadLunge                                 LungeExerciseName = 0
	LungeExerciseNameLungeMatrix                                   LungeExerciseName = 1
	LungeExerciseNameWeightedLungeMatrix                           LungeExerciseName = 2
	LungeExerciseNameAlternatingBarbellForwardLunge                LungeExerciseName = 3
	LungeExerciseNameAlternatingDumbbellLungeWithReach             LungeExerciseName = 4
	LungeExerciseNameBackFootElevatedDumbbellSplitSquat            LungeExerciseName = 5
	LungeExerciseNameBarbellBoxLunge                               LungeExerciseName = 6
	LungeExerciseNameBarbellBulgarianSplitSquat                    LungeExerciseName = 7
	LungeExerciseNameBarbellCrossoverLunge                         LungeExerciseName = 8
	LungeExerciseNameBarbellFrontSplitSquat                        LungeExerciseName = 9
	LungeExerciseNameBarbellLunge                                  LungeExerciseName = 10
	LungeExerciseNameBarbellReverseLunge                           LungeExerciseName = 11
	LungeExerciseNameBarbellSideLunge                              LungeExerciseName = 12
	LungeExerciseNameBarbellSplitSquat                             LungeExerciseName = 13
	LungeExerciseNameCoreControlRearLunge                          LungeExerciseName = 14
	LungeExerciseNameDiagonalLunge                                 LungeExerciseName = 15
	LungeExerciseNameDropLunge                                     LungeExerciseName = 16
	LungeExerciseNameDumbbellBoxLunge                              LungeExerciseName = 17
	LungeExerciseNameDumbbellBulgarianSplitSquat                   LungeExerciseName = 18
	LungeExerciseNameDumbbellCrossoverLunge                        LungeExerciseName = 19
	LungeExerciseNameDumbbellDiagonalLunge                         LungeExerciseName = 20
	LungeExerciseNameDumbbellLunge                                 LungeExerciseName = 21
	LungeExerciseNameDumbbellLungeAndRotation                      LungeExerciseName = 22
	LungeExerciseNameDumbbellOverheadBulgarianSplitSquat           LungeExerciseName = 23
	LungeExerciseNameDumbbellReverseLungeToHighKneeAndPress        LungeExerciseName = 24
	LungeExerciseNameDumbbellSideLunge                             LungeExerciseName = 25
	LungeExerciseNameElevatedFrontFootBarbellSplitSquat            LungeExerciseName = 26
	LungeExerciseNameFrontFootElevatedDumbbellSplitSquat           LungeExerciseName = 27
	LungeExerciseNameGunslingerLunge                               LungeExerciseName = 28
	LungeExerciseNameLawnmowerLunge                                LungeExerciseName = 29
	LungeExerciseNameLowLungeWithIsometricAdduction                LungeExerciseName = 30
	LungeExerciseNameLowSideToSideLunge                            LungeExerciseName = 31
	LungeExerciseNameLunge                                         LungeExerciseName = 32
	LungeExerciseNameWeightedLunge                                 LungeExerciseName = 33
	LungeExerciseNameLungeWithArmReach                             LungeExerciseName = 34
	LungeExerciseNameLungeWithDiagonalReach                        LungeExerciseName = 35
	LungeExerciseNameLungeWithSideBend                             LungeExerciseName = 36
	LungeExerciseNameOffsetDumbbellLunge                           LungeExerciseName = 37
	LungeExerciseNameOffsetDumbbellReverseLunge                    LungeExerciseName = 38
	LungeExerciseNameOverheadBulgarianSplitSquat                   LungeExerciseName = 39
	LungeExerciseNameOverheadDumbbellReverseLunge                  LungeExerciseName = 40
	LungeExerciseNameOverheadDumbbellSplitSquat                    LungeExerciseName = 41
	LungeExerciseNameOverheadLungeWithRotation                     LungeExerciseName = 42
	LungeExerciseNameReverseBarbellBoxLunge                        LungeExerciseName = 43
	LungeExerciseNameReverseBoxLunge                               LungeExerciseName = 44
	LungeExerciseNameReverseDumbbellBoxLunge                       LungeExerciseName = 45
	LungeExerciseNameReverseDumbbellCrossoverLunge                 LungeExerciseName = 46
	LungeExerciseNameReverseDumbbellDiagonalLunge                  LungeExerciseName = 47
	LungeExerciseNameReverseLungeWithReachBack                     LungeExerciseName = 48
	LungeExerciseNameWeightedReverseLungeWithReachBack             LungeExerciseName = 49
	LungeExerciseNameReverseLungeWithTwistAndOverheadReach         LungeExerciseName = 50
	LungeExerciseNameWeightedReverseLungeWithTwistAndOverheadReach LungeExerciseName = 51
	LungeExerciseNameReverseSlidingBoxLunge                        LungeExerciseName = 52
	LungeExerciseNameWeightedReverseSlidingBoxLunge                LungeExerciseName = 53
	LungeExerciseNameReverseSlidingLunge                           LungeExerciseName = 54
	LungeExerciseNameWeightedReverseSlidingLunge                   LungeExerciseName = 55
	LungeExerciseNameRunnersLungeToBalance                         LungeExerciseName = 56
	LungeExerciseNameWeightedRunnersLungeToBalance                 LungeExerciseName = 57
	LungeExerciseNameShiftingSideLunge                             LungeExerciseName = 58
	LungeExerciseNameSideAndCrossoverLunge                         LungeExerciseName = 59
	LungeExerciseNameWeightedSideAndCrossoverLunge                 LungeExerciseName = 60
	LungeExerciseNameSideLunge                                     LungeExerciseName = 61
	LungeExerciseNameWeightedSideLunge                             LungeExerciseName = 62
	LungeExerciseNameSideLungeAndPress                             LungeExerciseName = 63
	LungeExerciseNameSideLungeJumpOff                              LungeExerciseName = 64
	LungeExerciseNameSideLungeSweep                                LungeExerciseName = 65
	LungeExerciseNameWeightedSideLungeSweep                        LungeExerciseName = 66
	LungeExerciseNameSideLungeToCrossoverTap                       LungeExerciseName = 67
	LungeExerciseNameWeightedSideLungeToCrossoverTap               LungeExerciseName = 68
	LungeExerciseNameSideToSideLungeChops                          LungeExerciseName = 69
	LungeExerciseNameWeightedSideToSideLungeChops                  LungeExerciseName = 70
	LungeExerciseNameSiffJumpLunge                                 LungeExerciseName = 71
	LungeExerciseNameWeightedSiffJumpLunge                         LungeExerciseName = 72
	LungeExerciseNameSingleArmReverseLungeAndPress                 LungeExerciseName = 73
	LungeExerciseNameSlidingLateralLunge                           LungeExerciseName = 74
	LungeExerciseNameWeightedSlidingLateralLunge                   LungeExerciseName = 75
	LungeExerciseNameWalkingBarbellLunge                           LungeExerciseName = 76
	LungeExerciseNameWalkingDumbbellLunge                          LungeExerciseName = 77
	LungeExerciseNameWalkingLunge                                  LungeExerciseName = 78
	LungeExerciseNameWeightedWalkingLunge                          LungeExerciseName = 79
	LungeExerciseNameWideGripOverheadBarbellSplitSquat             LungeExerciseName = 80
	LungeExerciseNameInvalid                                       LungeExerciseName = 0xFFFF
)

// Manufacturer represents the manufacturer FIT type.
type Manufacturer uint16

const (
	ManufacturerGarmin                 Manufacturer = 1
	ManufacturerGarminFr405Antfs       Manufacturer = 2 // Do not use. Used by FR405 for ANTFS man id.
	ManufacturerZephyr                 Manufacturer = 3
	ManufacturerDayton                 Manufacturer = 4
	ManufacturerIdt                    Manufacturer = 5
	ManufacturerSrm                    Manufacturer = 6
	ManufacturerQuarq                  Manufacturer = 7
	ManufacturerIbike                  Manufacturer = 8
	ManufacturerSaris                  Manufacturer = 9
	ManufacturerSparkHk                Manufacturer = 10
	ManufacturerTanita                 Manufacturer = 11
	ManufacturerEchowell               Manufacturer = 12
	ManufacturerDynastreamOem          Manufacturer = 13
	ManufacturerNautilus               Manufacturer = 14
	ManufacturerDynastream             Manufacturer = 15
	ManufacturerTimex                  Manufacturer = 16
	ManufacturerMetrigear              Manufacturer = 17
	ManufacturerXelic                  Manufacturer = 18
	ManufacturerBeurer                 Manufacturer = 19
	ManufacturerCardiosport            Manufacturer = 20
	ManufacturerAAndD                  Manufacturer = 21
	ManufacturerHmm                    Manufacturer = 22
	ManufacturerSuunto                 Manufacturer = 23
	ManufacturerThitaElektronik        Manufacturer = 24
	ManufacturerGpulse                 Manufacturer = 25
	ManufacturerCleanMobile            Manufacturer = 26
	ManufacturerPedalBrain             Manufacturer = 27
	ManufacturerPeaksware              Manufacturer = 28
	ManufacturerSaxonar                Manufacturer = 29
	ManufacturerLemondFitness          Manufacturer = 30
	ManufacturerDexcom                 Manufacturer = 31
	ManufacturerWahooFitness           Manufacturer = 32
	ManufacturerOctaneFitness          Manufacturer = 33
	ManufacturerArchinoetics           Manufacturer = 34
	ManufacturerTheHurtBox             Manufacturer = 35
	ManufacturerCitizenSystems         Manufacturer = 36
	ManufacturerMagellan               Manufacturer = 37
	ManufacturerOsynce                 Manufacturer = 38
	ManufacturerHolux                  Manufacturer = 39
	ManufacturerConcept2               Manufacturer = 40
	ManufacturerShimano                Manufacturer = 41
	ManufacturerOneGiantLeap           Manufacturer = 42
	ManufacturerAceSensor              Manufacturer = 43
	ManufacturerBrimBrothers           Manufacturer = 44
	ManufacturerXplova                 Manufacturer = 45
	ManufacturerPerceptionDigital      Manufacturer = 46
	ManufacturerBf1systems             Manufacturer = 47
	ManufacturerPioneer                Manufacturer = 48
	ManufacturerSpantec                Manufacturer = 49
	ManufacturerMetalogics             Manufacturer = 50
	Manufacturer4iiiis                 Manufacturer = 51
	ManufacturerSeikoEpson             Manufacturer = 52
	ManufacturerSeikoEpsonOem          Manufacturer = 53
	ManufacturerIforPowell             Manufacturer = 54
	ManufacturerMaxwellGuider          Manufacturer = 55
	ManufacturerStarTrac               Manufacturer = 56
	ManufacturerBreakaway              Manufacturer = 57
	ManufacturerAlatechTechnologyLtd   Manufacturer = 58
	ManufacturerMioTechnologyEurope    Manufacturer = 59
	ManufacturerRotor                  Manufacturer = 60
	ManufacturerGeonaute               Manufacturer = 61
	ManufacturerIdBike                 Manufacturer = 62
	ManufacturerSpecialized            Manufacturer = 63
	ManufacturerWtek                   Manufacturer = 64
	ManufacturerPhysicalEnterprises    Manufacturer = 65
	ManufacturerNorthPoleEngineering   Manufacturer = 66
	ManufacturerBkool                  Manufacturer = 67
	ManufacturerCateye                 Manufacturer = 68
	ManufacturerStagesCycling          Manufacturer = 69
	ManufacturerSigmasport             Manufacturer = 70
	ManufacturerTomtom                 Manufacturer = 71
	ManufacturerPeripedal              Manufacturer = 72
	ManufacturerWattbike               Manufacturer = 73
	ManufacturerMoxy                   Manufacturer = 76
	ManufacturerCiclosport             Manufacturer = 77
	ManufacturerPowerbahn              Manufacturer = 78
	ManufacturerAcornProjectsAps       Manufacturer = 79
	ManufacturerLifebeam               Manufacturer = 80
	ManufacturerBontrager              Manufacturer = 81
	ManufacturerWellgo                 Manufacturer = 82
	ManufacturerScosche                Manufacturer = 83
	ManufacturerMagura                 Manufacturer = 84
	ManufacturerWoodway                Manufacturer = 85
	ManufacturerElite                  Manufacturer = 86
	ManufacturerNielsenKellerman       Manufacturer = 87
	ManufacturerDkCity                 Manufacturer = 88
	ManufacturerTacx                   Manufacturer = 89
	ManufacturerDirectionTechnology    Manufacturer = 90
	ManufacturerMagtonic               Manufacturer = 91
	Manufacturer1partcarbon            Manufacturer = 92
	ManufacturerInsideRideTechnologies Manufacturer = 93
	ManufacturerSoundOfMotion          Manufacturer = 94
	ManufacturerStryd                  Manufacturer = 95
	ManufacturerIcg                    Manufacturer = 96 // Indoorcycling Group
	ManufacturerMiPulse                Manufacturer = 97
	ManufacturerBsxAthletics           Manufacturer = 98
	ManufacturerLook                   Manufacturer = 99
	ManufacturerCampagnoloSrl          Manufacturer = 100
	ManufacturerBodyBikeSmart          Manufacturer = 101
	ManufacturerPraxisworks            Manufacturer = 102
	ManufacturerLimitsTechnology       Manufacturer = 103 // Limits Technology Ltd.
	ManufacturerTopactionTechnology    Manufacturer = 104 // TopAction Technology Inc.
	ManufacturerCosinuss               Manufacturer = 105
	ManufacturerFitcare                Manufacturer = 106
	ManufacturerMagene                 Manufacturer = 107
	ManufacturerGiantManufacturingCo   Manufacturer = 108
	ManufacturerTigrasport             Manufacturer = 109 // Tigrasport
	ManufacturerSalutron               Manufacturer = 110
	ManufacturerTechnogym              Manufacturer = 111
	ManufacturerBrytonSensors          Manufacturer = 112
	ManufacturerLatitudeLimited        Manufacturer = 113
	ManufacturerSoaringTechnology      Manufacturer = 114
	ManufacturerIgpsport               Manufacturer = 115
	ManufacturerThinkrider             Manufacturer = 116
	ManufacturerGopherSport            Manufacturer = 117
	ManufacturerWaterrower             Manufacturer = 118
	ManufacturerOrangetheory           Manufacturer = 119
	ManufacturerInpeak                 Manufacturer = 120
	ManufacturerKinetic                Manufacturer = 121
	ManufacturerJohnsonHealthTech      Manufacturer = 122
	ManufacturerPolarElectro           Manufacturer = 123
	ManufacturerSeesense               Manufacturer = 124
	ManufacturerNciTechnology          Manufacturer = 125
	ManufacturerIqsquare               Manufacturer = 126
	ManufacturerLeomo                  Manufacturer = 127
	ManufacturerIfitCom                Manufacturer = 128
	ManufacturerCorosByte              Manufacturer = 129
	ManufacturerVersaDesign            Manufacturer = 130
	ManufacturerChileaf                Manufacturer = 131
	ManufacturerCycplus                Manufacturer = 132
	ManufacturerGravaaByte             Manufacturer = 133
	ManufacturerSigeyi                 Manufacturer = 134
	ManufacturerCoospo                 Manufacturer = 135
	ManufacturerGeoid                  Manufacturer = 136
	ManufacturerBosch                  Manufacturer = 137
	ManufacturerKyto                   Manufacturer = 138
	ManufacturerKineticSports          Manufacturer = 139
	ManufacturerDecathlonByte          Manufacturer = 140
	ManufacturerTqSystems              Manufacturer = 141
	ManufacturerTagHeuer               Manufacturer = 142
	ManufacturerKeiserFitness          Manufacturer = 143
	ManufacturerZwiftByte              Manufacturer = 144
	ManufacturerPorscheEp              Manufacturer = 145
	ManufacturerBlackbird              Manufacturer = 146
	ManufacturerMeilanByte             Manufacturer = 147
	ManufacturerEzon                   Manufacturer = 148
	ManufacturerDevelopment            Manufacturer = 255
	ManufacturerHealthandlife          Manufacturer = 257
	ManufacturerLezyne                 Manufacturer = 258
	ManufacturerScribeLabs             Manufacturer = 259
	ManufacturerZwift                  Manufacturer = 260
	ManufacturerWatteam                Manufacturer = 261
	ManufacturerRecon                  Manufacturer = 262
	ManufacturerFaveroElectronics      Manufacturer = 263
	ManufacturerDynovelo               Manufacturer = 264
	ManufacturerStrava                 Manufacturer = 265
	ManufacturerPrecor                 Manufacturer = 266 // Amer Sports
	ManufacturerBryton                 Manufacturer = 267
	ManufacturerSram                   Manufacturer = 268
	ManufacturerNavman                 Manufacturer = 269 // MiTAC Global Corporation (Mio Technology)
	ManufacturerCobi                   Manufacturer = 270 // COBI GmbH
	ManufacturerSpivi                  Manufacturer = 271
	ManufacturerMioMagellan            Manufacturer = 272
	ManufacturerEvesports              Manufacturer = 273
	ManufacturerSensitivusGauge        Manufacturer = 274
	ManufacturerPodoon                 Manufacturer = 275
	ManufacturerLifeTimeFitness        Manufacturer = 276
	ManufacturerFalcoEMotors           Manufacturer = 277 // Falco eMotors Inc.
	ManufacturerMinoura                Manufacturer = 278
	ManufacturerCycliq                 Manufacturer = 279
	ManufacturerLuxottica              Manufacturer = 280
	ManufacturerTrainerRoad            Manufacturer = 281
	ManufacturerTheSufferfest          Manufacturer = 282
	ManufacturerFullspeedahead         Manufacturer = 283
	ManufacturerVirtualtraining        Manufacturer = 284
	ManufacturerFeedbacksports         Manufacturer = 285
	ManufacturerOmata                  Manufacturer = 286
	ManufacturerVdo                    Manufacturer = 287
	ManufacturerMagneticdays           Manufacturer = 288
	ManufacturerHammerhead             Manufacturer = 289
	ManufacturerKineticByKurt          Manufacturer = 290
	ManufacturerShapelog               Manufacturer = 291
	ManufacturerDabuziduo              Manufacturer = 292
	ManufacturerJetblack               Manufacturer = 293
	ManufacturerCoros                  Manufacturer = 294
	ManufacturerVirtugo                Manufacturer = 295
	ManufacturerVelosense              Manufacturer = 296
	ManufacturerCycligentinc           Manufacturer = 297
	ManufacturerTrailforks             Manufacturer = 298
	ManufacturerMahleEbikemotion       Manufacturer = 299
	ManufacturerNurvv                  Manufacturer = 300
	ManufacturerMicroprogram           Manufacturer = 301
	ManufacturerZone5cloud             Manufacturer = 302
	ManufacturerGreenteg               Manufacturer = 303
	ManufacturerYamahaMotors           Manufacturer = 304
	ManufacturerWhoop                  Manufacturer = 305
	ManufacturerGravaa                 Manufacturer = 306
	ManufacturerOnelap                 Manufacturer = 307
	ManufacturerMonarkExercise         Manufacturer = 308
	ManufacturerForm                   Manufacturer = 309
	ManufacturerDecathlon              Manufacturer = 310
	ManufacturerSyncros                Manufacturer = 311
	ManufacturerHeatup                 Manufacturer = 312
	ManufacturerCannondale             Manufacturer = 313
	ManufacturerTrueFitness            Manufacturer = 314
	ManufacturerRGTCycling             Manufacturer = 315
	ManufacturerVasa                   Manufacturer = 316
	ManufacturerRaceRepublic           Manufacturer = 317
	ManufacturerFazua                  Manufacturer = 318
	ManufacturerOrekaTraining          Manufacturer = 319
	ManufacturerLsec                   Manufacturer = 320 // Lishun Electric & Communication
	ManufacturerLululemonStudio        Manufacturer = 321
	ManufacturerShanyue                Manufacturer = 322
	ManufacturerSpinningMda            Manufacturer = 323
	ManufacturerHilldating             Manufacturer = 324
	ManufacturerActigraphcorp          Manufacturer = 5759
	ManufacturerInvalid                Manufacturer = 0xFFFF
)

// MaxMetCategory represents the max_met_category FIT type.
type MaxMetCategory byte

const (
	MaxMetCategoryGeneric MaxMetCategory = 0
	MaxMetCategoryCycling MaxMetCategory = 1
	MaxMetCategoryInvalid MaxMetCategory = 0xFF
)

// MaxMetHeartRateSource represents the max_met_heart_rate_source FIT type.
type MaxMetHeartRateSource byte

const (
	MaxMetHeartRateSourceWhr     MaxMetHeartRateSource = 0 // Wrist Heart Rate Monitor
	MaxMetHeartRateSourceHrm     MaxMetHeartRateSource = 1 // Chest Strap Heart Rate Monitor
	MaxMetHeartRateSourceInvalid MaxMetHeartRateSource = 0xFF
)

// MaxMetSpeedSource represents the max_met_speed_source FIT type.
type MaxMetSpeedSource byte

const (
	MaxMetSpeedSourceOnboardGps   MaxMetSpeedSource = 0
	MaxMetSpeedSourceConnectedGps MaxMetSpeedSource = 1
	MaxMetSpeedSourceCadence      MaxMetSpeedSource = 2
	MaxMetSpeedSourceInvalid      MaxMetSpeedSource = 0xFF
)

// MesgCount represents the mesg_count FIT type.
type MesgCount byte

const (
	MesgCountNumPerFile     MesgCount = 0
	MesgCountMaxPerFile     MesgCount = 1
	MesgCountMaxPerFileType MesgCount = 2
	MesgCountInvalid        MesgCount = 0xFF
)

// MesgNum represents the mesg_num FIT type.
type MesgNum uint16

const (
	MesgNumFileId                      MesgNum = 0
	MesgNumCapabilities                MesgNum = 1
	MesgNumDeviceSettings              MesgNum = 2
	MesgNumUserProfile                 MesgNum = 3
	MesgNumHrmProfile                  MesgNum = 4
	MesgNumSdmProfile                  MesgNum = 5
	MesgNumBikeProfile                 MesgNum = 6
	MesgNumZonesTarget                 MesgNum = 7
	MesgNumHrZone                      MesgNum = 8
	MesgNumPowerZone                   MesgNum = 9
	MesgNumMetZone                     MesgNum = 10
	MesgNumSport                       MesgNum = 12
	MesgNumGoal                        MesgNum = 15
	MesgNumSession                     MesgNum = 18
	MesgNumLap                         MesgNum = 19
	MesgNumRecord                      MesgNum = 20
	MesgNumEvent                       MesgNum = 21
	MesgNumDeviceInfo                  MesgNum = 23
	MesgNumWorkout                     MesgNum = 26
	MesgNumWorkoutStep                 MesgNum = 27
	MesgNumSchedule                    MesgNum = 28
	MesgNumWeightScale                 MesgNum = 30
	MesgNumCourse                      MesgNum = 31
	MesgNumCoursePoint                 MesgNum = 32
	MesgNumTotals                      MesgNum = 33
	MesgNumActivity                    MesgNum = 34
	MesgNumSoftware                    MesgNum = 35
	MesgNumFileCapabilities            MesgNum = 37
	MesgNumMesgCapabilities            MesgNum = 38
	MesgNumFieldCapabilities           MesgNum = 39
	MesgNumFileCreator                 MesgNum = 49
	MesgNumBloodPressure               MesgNum = 51
	MesgNumSpeedZone                   MesgNum = 53
	MesgNumMonitoring                  MesgNum = 55
	MesgNumTrainingFile                MesgNum = 72
	MesgNumHrv                         MesgNum = 78
	MesgNumAntRx                       MesgNum = 80
	MesgNumAntTx                       MesgNum = 81
	MesgNumAntChannelId                MesgNum = 82
	MesgNumLength                      MesgNum = 101
	MesgNumMonitoringInfo              MesgNum = 103
	MesgNumPad                         MesgNum = 105
	MesgNumSlaveDevice                 MesgNum = 106
	MesgNumConnectivity                MesgNum = 127
	MesgNumWeatherConditions           MesgNum = 128
	MesgNumWeatherAlert                MesgNum = 129
	MesgNumCadenceZone                 MesgNum = 131
	MesgNumHr                          MesgNum = 132
	MesgNumSegmentLap                  MesgNum = 142
	MesgNumMemoGlob                    MesgNum = 145
	MesgNumSegmentId                   MesgNum = 148
	MesgNumSegmentLeaderboardEntry     MesgNum = 149
	MesgNumSegmentPoint                MesgNum = 150
	MesgNumSegmentFile                 MesgNum = 151
	MesgNumWorkoutSession              MesgNum = 158
	MesgNumWatchfaceSettings           MesgNum = 159
	MesgNumGpsMetadata                 MesgNum = 160
	MesgNumCameraEvent                 MesgNum = 161
	MesgNumTimestampCorrelation        MesgNum = 162
	MesgNumGyroscopeData               MesgNum = 164
	MesgNumAccelerometerData           MesgNum = 165
	MesgNumThreeDSensorCalibration     MesgNum = 167
	MesgNumVideoFrame                  MesgNum = 169
	MesgNumObdiiData                   MesgNum = 174
	MesgNumNmeaSentence                MesgNum = 177
	MesgNumAviationAttitude            MesgNum = 178
	MesgNumVideo                       MesgNum = 184
	MesgNumVideoTitle                  MesgNum = 185
	MesgNumVideoDescription            MesgNum = 186
	MesgNumVideoClip                   MesgNum = 187
	MesgNumOhrSettings                 MesgNum = 188
	MesgNumExdScreenConfiguration      MesgNum = 200
	MesgNumExdDataFieldConfiguration   MesgNum = 201
	MesgNumExdDataConceptConfiguration MesgNum = 202
	MesgNumFieldDescription            MesgNum = 206
	MesgNumDeveloperDataId             MesgNum = 207
	MesgNumMagnetometerData            MesgNum = 208
	MesgNumBarometerData               MesgNum = 209
	MesgNumOneDSensorCalibration       MesgNum = 210
	MesgNumMonitoringHrData            MesgNum = 211
	MesgNumTimeInZone                  MesgNum = 216
	MesgNumSet                         MesgNum = 225
	MesgNumStressLevel                 MesgNum = 227
	MesgNumMaxMetData                  MesgNum = 229
	MesgNumDiveSettings                MesgNum = 258
	MesgNumDiveGas                     MesgNum = 259
	MesgNumDiveAlarm                   MesgNum = 262
	MesgNumExerciseTitle               MesgNum = 264
	MesgNumDiveSummary                 MesgNum = 268
	MesgNumSpo2Data                    MesgNum = 269
	MesgNumSleepLevel                  MesgNum = 275
	MesgNumJump                        MesgNum = 285
	MesgNumBeatIntervals               MesgNum = 290
	MesgNumRespirationRate             MesgNum = 297
	MesgNumSplit                       MesgNum = 312
	MesgNumClimbPro                    MesgNum = 317
	MesgNumTankUpdate                  MesgNum = 319
	MesgNumTankSummary                 MesgNum = 323
	MesgNumSleepAssessment             MesgNum = 346
	MesgNumHrvStatusSummary            MesgNum = 370
	MesgNumHrvValue                    MesgNum = 371
	MesgNumDeviceAuxBatteryInfo        MesgNum = 375
	MesgNumDiveApneaAlarm              MesgNum = 393
	MesgNumMfgRangeMin                 MesgNum = 0xFF00 // 0xFF00 - 0xFFFE reserved for manufacturer specific messages
	MesgNumMfgRangeMax                 MesgNum = 0xFFFE // 0xFF00 - 0xFFFE reserved for manufacturer specific messages
	MesgNumInvalid                     MesgNum = 0xFFFF
)

// MessageIndex represents the message_index FIT type.
type MessageIndex uint16

const (
	MessageIndexSelected MessageIndex = 0x8000 // message is selected if set
	MessageIndexReserved MessageIndex = 0x7000 // reserved (default 0)
	MessageIndexMask     MessageIndex = 0x0FFF // index
	MessageIndexInvalid  MessageIndex = 0xFFFF
)

// NoFlyTimeMode represents the no_fly_time_mode FIT type.
type NoFlyTimeMode byte

const (
	NoFlyTimeModeStandard    NoFlyTimeMode = 0 // Standard Diver Alert Network no-fly guidance
	NoFlyTimeModeFlat24Hours NoFlyTimeMode = 1 // Flat 24 hour no-fly guidance
	NoFlyTimeModeInvalid     NoFlyTimeMode = 0xFF
)

// OlympicLiftExerciseName represents the olympic_lift_exercise_name FIT type.
type OlympicLiftExerciseName uint16

const (
	OlympicLiftExerciseNameBarbellHangPowerClean      OlympicLiftExerciseName = 0
	OlympicLiftExerciseNameBarbellHangSquatClean      OlympicLiftExerciseName = 1
	OlympicLiftExerciseNameBarbellPowerClean          OlympicLiftExerciseName = 2
	OlympicLiftExerciseNameBarbellPowerSnatch         OlympicLiftExerciseName = 3
	OlympicLiftExerciseNameBarbellSquatClean          OlympicLiftExerciseName = 4
	OlympicLiftExerciseNameCleanAndJerk               OlympicLiftExerciseName = 5
	OlympicLiftExerciseNameBarbellHangPowerSnatch     OlympicLiftExerciseName = 6
	OlympicLiftExerciseNameBarbellHangPull            OlympicLiftExerciseName = 7
	OlympicLiftExerciseNameBarbellHighPull            OlympicLiftExerciseName = 8
	OlympicLiftExerciseNameBarbellSnatch              OlympicLiftExerciseName = 9
	OlympicLiftExerciseNameBarbellSplitJerk           OlympicLiftExerciseName = 10
	OlympicLiftExerciseNameClean                      OlympicLiftExerciseName = 11
	OlympicLiftExerciseNameDumbbellClean              OlympicLiftExerciseName = 12
	OlympicLiftExerciseNameDumbbellHangPull           OlympicLiftExerciseName = 13
	OlympicLiftExerciseNameOneHandDumbbellSplitSnatch OlympicLiftExerciseName = 14
	OlympicLiftExerciseNamePushJerk                   OlympicLiftExerciseName = 15
	OlympicLiftExerciseNameSingleArmDumbbellSnatch    OlympicLiftExerciseName = 16
	OlympicLiftExerciseNameSingleArmHangSnatch        OlympicLiftExerciseName = 17
	OlympicLiftExerciseNameSingleArmKettlebellSnatch  OlympicLiftExerciseName = 18
	OlympicLiftExerciseNameSplitJerk                  OlympicLiftExerciseName = 19
	OlympicLiftExerciseNameSquatCleanAndJerk          OlympicLiftExerciseName = 20
	OlympicLiftExerciseNameInvalid                    OlympicLiftExerciseName = 0xFFFF
)

// PlankExerciseName represents the plank_exercise_name FIT type.
type PlankExerciseName uint16

const (
	PlankExerciseName45DegreePlank                                    PlankExerciseName = 0
	PlankExerciseNameWeighted45DegreePlank                            PlankExerciseName = 1
	PlankExerciseName90DegreeStaticHold                               PlankExerciseName = 2
	PlankExerciseNameWeighted90DegreeStaticHold                       PlankExerciseName = 3
	PlankExerciseNameBearCrawl                                        PlankExerciseName = 4
	PlankExerciseNameWeightedBearCrawl                                PlankExerciseName = 5
	PlankExerciseNameCrossBodyMountainClimber                         PlankExerciseName = 6
	PlankExerciseNameWeightedCrossBodyMountainClimber                 PlankExerciseName = 7
	PlankExerciseNameElbowPlankPikeJacks                              PlankExerciseName = 8
	PlankExerciseNameWeightedElbowPlankPikeJacks                      PlankExerciseName = 9
	PlankExerciseNameElevatedFeetPlank                                PlankExerciseName = 10
	PlankExerciseNameWeightedElevatedFeetPlank                        PlankExerciseName = 11
	PlankExerciseNameElevatorAbs                                      PlankExerciseName = 12
	PlankExerciseNameWeightedElevatorAbs                              PlankExerciseName = 13
	PlankExerciseNameExtendedPlank                                    PlankExerciseName = 14
	PlankExerciseNameWeightedExtendedPlank                            PlankExerciseName = 15
	PlankExerciseNameFullPlankPasseTwist                              PlankExerciseName = 16
	PlankExerciseNameWeightedFullPlankPasseTwist                      PlankExerciseName = 17
	PlankExerciseNameInchingElbowPlank                                PlankExerciseName = 18
	PlankExerciseNameWeightedInchingElbowPlank                        PlankExerciseName = 19
	PlankExerciseNameInchwormToSidePlank                              PlankExerciseName = 20
	PlankExerciseNameWeightedInchwormToSidePlank                      PlankExerciseName = 21
	PlankExerciseNameKneelingPlank                                    PlankExerciseName = 22
	PlankExerciseNameWeightedKneelingPlank                            PlankExerciseName = 23
	PlankExerciseNameKneelingSidePlankWithLegLift                     PlankExerciseName = 24
	PlankExerciseNameWeightedKneelingSidePlankWithLegLift             PlankExerciseName = 25
	PlankExerciseNameLateralRoll                                      PlankExerciseName = 26
	PlankExerciseNameWeightedLateralRoll                              PlankExerciseName = 27
	PlankExerciseNameLyingReversePlank                                PlankExerciseName = 28
	PlankExerciseNameWeightedLyingReversePlank                        PlankExerciseName = 29
	PlankExerciseNameMedicineBallMountainClimber                      PlankExerciseName = 30
	PlankExerciseNameWeightedMedicineBallMountainClimber              PlankExerciseName = 31
	PlankExerciseNameModifiedMountainClimberAndExtension              PlankExerciseName = 32
	PlankExerciseNameWeightedModifiedMountainClimberAndExtension      PlankExerciseName = 33
	PlankExerciseNameMountainClimber                                  PlankExerciseName = 34
	PlankExerciseNameWeightedMountainClimber                          PlankExerciseName = 35
	PlankExerciseNameMountainClimberOnSlidingDiscs                    PlankExerciseName = 36
	PlankExerciseNameWeightedMountainClimberOnSlidingDiscs            PlankExerciseName = 37
	PlankExerciseNameMountainClimberWithFeetOnBosuBall                PlankExerciseName = 38
	PlankExerciseNameWeightedMountainClimberWithFeetOnBosuBall        PlankExerciseName = 39
	PlankExerciseNameMountainClimberWithHandsOnBench                  PlankExerciseName = 40
	PlankExerciseNameMountainClimberWithHandsOnSwissBall              PlankExerciseName = 41
	PlankExerciseNameWeightedMountainClimberWithHandsOnSwissBall      PlankExerciseName = 42
	PlankExerciseNamePlank                                            PlankExerciseName = 43
	PlankExerciseNamePlankJacksWithFeetOnSlidingDiscs                 PlankExerciseName = 44
	PlankExerciseNameWeightedPlankJacksWithFeetOnSlidingDiscs         PlankExerciseName = 45
	PlankExerciseNamePlankKneeTwist                                   PlankExerciseName = 46
	PlankExerciseNameWeightedPlankKneeTwist                           PlankExerciseName = 47
	PlankExerciseNamePlankPikeJumps                                   PlankExerciseName = 48
	PlankExerciseNameWeightedPlankPikeJumps                           PlankExerciseName = 49
	PlankExerciseNamePlankPikes                                       PlankExerciseName = 50
	PlankExerciseNameWeightedPlankPikes                               PlankExerciseName = 51
	PlankExerciseNamePlankToStandUp                                   PlankExerciseName = 52
	PlankExerciseNameWeightedPlankToStandUp                           PlankExerciseName = 53
	PlankExerciseNamePlankWithArmRaise                                PlankExerciseName = 54
	PlankExerciseNameWeightedPlankWithArmRaise                        PlankExerciseName = 55
	PlankExerciseNamePlankWithKneeToElbow                             PlankExerciseName = 56
	PlankExerciseNameWeightedPlankWithKneeToElbow                     PlankExerciseName = 57
	PlankExerciseNamePlankWithObliqueCrunch                           PlankExerciseName = 58
	PlankExerciseNameWeightedPlankWithObliqueCrunch                   PlankExerciseName = 59
	PlankExerciseNamePlyometricSidePlank                              PlankExerciseName = 60
	PlankExerciseNameWeightedPlyometricSidePlank                      PlankExerciseName = 61
	PlankExerciseNameRollingSidePlank                                 PlankExerciseName = 62
	PlankExerciseNameWeightedRollingSidePlank                         PlankExerciseName = 63
	PlankExerciseNameSideKickPlank                                    PlankExerciseName = 64
	PlankExerciseNameWeightedSideKickPlank                            PlankExerciseName = 65
	PlankExerciseNameSidePlank                                        PlankExerciseName = 66
	PlankExerciseNameWeightedSidePlank                                PlankExerciseName = 67
	PlankExerciseNameSidePlankAndRow                                  PlankExerciseName = 68
	PlankExerciseNameWeightedSidePlankAndRow                          PlankExerciseName = 69
	PlankExerciseNameSidePlankLift                                    PlankExerciseName = 70
	PlankExerciseNameWeightedSidePlankLift                            PlankExerciseName = 71
	PlankExerciseNameSidePlankWithElbowOnBosuBall                     PlankExerciseName = 72
	PlankExerciseNameWeightedSidePlankWithElbowOnBosuBall             PlankExerciseName = 73
	PlankExerciseNameSidePlankWithFeetOnBench                         PlankExerciseName = 74
	PlankExerciseNameWeightedSidePlankWithFeetOnBench                 PlankExerciseName = 75
	PlankExerciseNameSidePlankWithKneeCircle                          PlankExerciseName = 76
	PlankExerciseNameWeightedSidePlankWithKneeCircle                  PlankExerciseName = 77
	PlankExerciseNameSidePlankWithKneeTuck                            PlankExerciseName = 78
	PlankExerciseNameWeightedSidePlankWithKneeTuck                    PlankExerciseName = 79
	PlankExerciseNameSidePlankWithLegLift                             PlankExerciseName = 80
	PlankExerciseNameWeightedSidePlankWithLegLift                     PlankExerciseName = 81
	PlankExerciseNameSidePlankWithReachUnder                          PlankExerciseName = 82
	PlankExerciseNameWeightedSidePlankWithReachUnder                  PlankExerciseName = 83
	PlankExerciseNameSingleLegElevatedFeetPlank                       PlankExerciseName = 84
	PlankExerciseNameWeightedSingleLegElevatedFeetPlank               PlankExerciseName = 85
	PlankExerciseNameSingleLegFlexAndExtend                           PlankExerciseName = 86
	PlankExerciseNameWeightedSingleLegFlexAndExtend                   PlankExerciseName = 87
	PlankExerciseNameSingleLegSidePlank                               PlankExerciseName = 88
	PlankExerciseNameWeightedSingleLegSidePlank                       PlankExerciseName = 89
	PlankExerciseNameSpidermanPlank                                   PlankExerciseName = 90
	PlankExerciseNameWeightedSpidermanPlank                           PlankExerciseName = 91
	PlankExerciseNameStraightArmPlank                                 PlankExerciseName = 92
	PlankExerciseNameWeightedStraightArmPlank                         PlankExerciseName = 93
	PlankExerciseNameStraightArmPlankWithShoulderTouch                PlankExerciseName = 94
	PlankExerciseNameWeightedStraightArmPlankWithShoulderTouch        PlankExerciseName = 95
	PlankExerciseNameSwissBallPlank                                   PlankExerciseName = 96
	PlankExerciseNameWeightedSwissBallPlank                           PlankExerciseName = 97
	PlankExerciseNameSwissBallPlankLegLift                            PlankExerciseName = 98
	PlankExerciseNameWeightedSwissBallPlankLegLift                    PlankExerciseName = 99
	PlankExerciseNameSwissBallPlankLegLiftAndHold                     PlankExerciseName = 100
	PlankExerciseNameSwissBallPlankWithFeetOnBench                    PlankExerciseName = 101
	PlankExerciseNameWeightedSwissBallPlankWithFeetOnBench            PlankExerciseName = 102
	PlankExerciseNameSwissBallProneJackknife                          PlankExerciseName = 103
	PlankExerciseNameWeightedSwissBallProneJackknife                  PlankExerciseName = 104
	PlankExerciseNameSwissBallSidePlank                               PlankExerciseName = 105
	PlankExerciseNameWeightedSwissBallSidePlank                       PlankExerciseName = 106
	PlankExerciseNameThreeWayPlank                                    PlankExerciseName = 107
	PlankExerciseNameWeightedThreeWayPlank                            PlankExerciseName = 108
	PlankExerciseNameTowelPlankAndKneeIn                              PlankExerciseName = 109
	PlankExerciseNameWeightedTowelPlankAndKneeIn                      PlankExerciseName = 110
	PlankExerciseNameTStabilization                                   PlankExerciseName = 111
	PlankExerciseNameWeightedTStabilization                           PlankExerciseName = 112
	PlankExerciseNameTurkishGetUpToSidePlank                          PlankExerciseName = 113
	PlankExerciseNameWeightedTurkishGetUpToSidePlank                  PlankExerciseName = 114
	PlankExerciseNameTwoPointPlank                                    PlankExerciseName = 115
	PlankExerciseNameWeightedTwoPointPlank                            PlankExerciseName = 116
	PlankExerciseNameWeightedPlank                                    PlankExerciseName = 117
	PlankExerciseNameWideStancePlankWithDiagonalArmLift               PlankExerciseName = 118
	PlankExerciseNameWeightedWideStancePlankWithDiagonalArmLift       PlankExerciseName = 119
	PlankExerciseNameWideStancePlankWithDiagonalLegLift               PlankExerciseName = 120
	PlankExerciseNameWeightedWideStancePlankWithDiagonalLegLift       PlankExerciseName = 121
	PlankExerciseNameWideStancePlankWithLegLift                       PlankExerciseName = 122
	PlankExerciseNameWeightedWideStancePlankWithLegLift               PlankExerciseName = 123
	PlankExerciseNameWideStancePlankWithOppositeArmAndLegLift         PlankExerciseName = 124
	PlankExerciseNameWeightedMountainClimberWithHandsOnBench          PlankExerciseName = 125
	PlankExerciseNameWeightedSwissBallPlankLegLiftAndHold             PlankExerciseName = 126
	PlankExerciseNameWeightedWideStancePlankWithOppositeArmAndLegLift PlankExerciseName = 127
	PlankExerciseNamePlankWithFeetOnSwissBall                         PlankExerciseName = 128
	PlankExerciseNameSidePlankToPlankWithReachUnder                   PlankExerciseName = 129
	PlankExerciseNameBridgeWithGluteLowerLift                         PlankExerciseName = 130
	PlankExerciseNameBridgeOneLegBridge                               PlankExerciseName = 131
	PlankExerciseNamePlankWithArmVariations                           PlankExerciseName = 132
	PlankExerciseNamePlankWithLegLift                                 PlankExerciseName = 133
	PlankExerciseNameReversePlankWithLegPull                          PlankExerciseName = 134
	PlankExerciseNameInvalid                                          PlankExerciseName = 0xFFFF
)

// PlyoExerciseName represents the plyo_exercise_name FIT type.
type PlyoExerciseName uint16

const (
	PlyoExerciseNameAlternatingJumpLunge                  PlyoExerciseName = 0
	PlyoExerciseNameWeightedAlternatingJumpLunge          PlyoExerciseName = 1
	PlyoExerciseNameBarbellJumpSquat                      PlyoExerciseName = 2
	PlyoExerciseNameBodyWeightJumpSquat                   PlyoExerciseName = 3
	PlyoExerciseNameWeightedJumpSquat                     PlyoExerciseName = 4
	PlyoExerciseNameCrossKneeStrike                       PlyoExerciseName = 5
	PlyoExerciseNameWeightedCrossKneeStrike               PlyoExerciseName = 6
	PlyoExerciseNameDepthJump                             PlyoExerciseName = 7
	PlyoExerciseNameWeightedDepthJump                     PlyoExerciseName = 8
	PlyoExerciseNameDumbbellJumpSquat                     PlyoExerciseName = 9
	PlyoExerciseNameDumbbellSplitJump                     PlyoExerciseName = 10
	PlyoExerciseNameFrontKneeStrike                       PlyoExerciseName = 11
	PlyoExerciseNameWeightedFrontKneeStrike               PlyoExerciseName = 12
	PlyoExerciseNameHighBoxJump                           PlyoExerciseName = 13
	PlyoExerciseNameWeightedHighBoxJump                   PlyoExerciseName = 14
	PlyoExerciseNameIsometricExplosiveBodyWeightJumpSquat PlyoExerciseName = 15
	PlyoExerciseNameWeightedIsometricExplosiveJumpSquat   PlyoExerciseName = 16
	PlyoExerciseNameLateralLeapAndHop                     PlyoExerciseName = 17
	PlyoExerciseNameWeightedLateralLeapAndHop             PlyoExerciseName = 18
	PlyoExerciseNameLateralPlyoSquats                     PlyoExerciseName = 19
	PlyoExerciseNameWeightedLateralPlyoSquats             PlyoExerciseName = 20
	PlyoExerciseNameLateralSlide                          PlyoExerciseName = 21
	PlyoExerciseNameWeightedLateralSlide                  PlyoExerciseName = 22
	PlyoExerciseNameMedicineBallOverheadThrows            PlyoExerciseName = 23
	PlyoExerciseNameMedicineBallSideThrow                 PlyoExerciseName = 24
	PlyoExerciseNameMedicineBallSlam                      PlyoExerciseName = 25
	PlyoExerciseNameSideToSideMedicineBallThrows          PlyoExerciseName = 26
	PlyoExerciseNameSideToSideShuffleJump                 PlyoExerciseName = 27
	PlyoExerciseNameWeightedSideToSideShuffleJump         PlyoExerciseName = 28
	PlyoExerciseNameSquatJumpOntoBox                      PlyoExerciseName = 29
	PlyoExerciseNameWeightedSquatJumpOntoBox              PlyoExerciseName = 30
	PlyoExerciseNameSquatJumpsInAndOut                    PlyoExerciseName = 31
	PlyoExerciseNameWeightedSquatJumpsInAndOut            PlyoExerciseName = 32
	PlyoExerciseNameInvalid                               PlyoExerciseName = 0xFFFF
)

// PowerPhaseType represents the power_phase_type FIT type.
type PowerPhaseType byte

const (
	PowerPhaseTypePowerPhaseStartAngle PowerPhaseType = 0
	PowerPhaseTypePowerPhaseEndAngle   PowerPhaseType = 1
	PowerPhaseTypePowerPhaseArcLength  PowerPhaseType = 2
	PowerPhaseTypePowerPhaseCenter     PowerPhaseType = 3
	PowerPhaseTypeInvalid              PowerPhaseType = 0xFF
)

// PullUpExerciseName represents the pull_up_exercise_name FIT type.
type PullUpExerciseName uint16

const (
	PullUpExerciseNameBandedPullUps                    PullUpExerciseName = 0
	PullUpExerciseName30DegreeLatPulldown              PullUpExerciseName = 1
	PullUpExerciseNameBandAssistedChinUp               PullUpExerciseName = 2
	PullUpExerciseNameCloseGripChinUp                  PullUpExerciseName = 3
	PullUpExerciseNameWeightedCloseGripChinUp          PullUpExerciseName = 4
	PullUpExerciseNameCloseGripLatPulldown             PullUpExerciseName = 5
	PullUpExerciseNameCrossoverChinUp                  PullUpExerciseName = 6
	PullUpExerciseNameWeightedCrossoverChinUp          PullUpExerciseName = 7
	PullUpExerciseNameEzBarPullover                    PullUpExerciseName = 8
	PullUpExerciseNameHangingHurdle                    PullUpExerciseName = 9
	PullUpExerciseNameWeightedHangingHurdle            PullUpExerciseName = 10
	PullUpExerciseNameKneelingLatPulldown              PullUpExerciseName = 11
	PullUpExerciseNameKneelingUnderhandGripLatPulldown PullUpExerciseName = 12
	PullUpExerciseNameLatPulldown                      PullUpExerciseName = 13
	PullUpExerciseNameMixedGripChinUp                  PullUpExerciseName = 14
	PullUpExerciseNameWeightedMixedGripChinUp          PullUpExerciseName = 15
	PullUpExerciseNameMixedGripPullUp                  PullUpExerciseName = 16
	PullUpExerciseNameWeightedMixedGripPullUp          PullUpExerciseName = 17
	PullUpExerciseNameReverseGripPulldown              PullUpExerciseName = 18
	PullUpExerciseNameStandingCablePullover            PullUpExerciseName = 19
	PullUpExerciseNameStraightArmPulldown              PullUpExerciseName = 20
	PullUpExerciseNameSwissBallEzBarPullover           PullUpExerciseName = 21
	PullUpExerciseNameTowelPullUp                      PullUpExerciseName = 22
	PullUpExerciseNameWeightedTowelPullUp              PullUpExerciseName = 23
	PullUpExerciseNameWeightedPullUp                   PullUpExerciseName = 24
	PullUpExerciseNameWideGripLatPulldown              PullUpExerciseName = 25
	PullUpExerciseNameWideGripPullUp                   PullUpExerciseName = 26
	PullUpExerciseNameWeightedWideGripPullUp           PullUpExerciseName = 27
	PullUpExerciseNameBurpeePullUp                     PullUpExerciseName = 28
	PullUpExerciseNameWeightedBurpeePullUp             PullUpExerciseName = 29
	PullUpExerciseNameJumpingPullUps                   PullUpExerciseName = 30
	PullUpExerciseNameWeightedJumpingPullUps           PullUpExerciseName = 31
	PullUpExerciseNameKippingPullUp                    PullUpExerciseName = 32
	PullUpExerciseNameWeightedKippingPullUp            PullUpExerciseName = 33
	PullUpExerciseNameLPullUp                          PullUpExerciseName = 34
	PullUpExerciseNameWeightedLPullUp                  PullUpExerciseName = 35
	PullUpExerciseNameSuspendedChinUp                  PullUpExerciseName = 36
	PullUpExerciseNameWeightedSuspendedChinUp          PullUpExerciseName = 37
	PullUpExerciseNamePullUp                           PullUpExerciseName = 38
	PullUpExerciseNameInvalid                          PullUpExerciseName = 0xFFFF
)

// PushUpExerciseName represents the push_up_exercise_name FIT type.
type PushUpExerciseName uint16

const (
	PushUpExerciseNameChestPressWithBand                         PushUpExerciseName = 0
	PushUpExerciseNameAlternatingStaggeredPushUp                 PushUpExerciseName = 1
	PushUpExerciseNameWeightedAlternatingStaggeredPushUp         PushUpExerciseName = 2
	PushUpExerciseNameAlternatingHandsMedicineBallPushUp         PushUpExerciseName = 3
	PushUpExerciseNameWeightedAlternatingHandsMedicineBallPushUp PushUpExerciseName = 4
	PushUpExerciseNameBosuBallPushUp                             PushUpExerciseName = 5
	PushUpExerciseNameWeightedBosuBallPushUp                     PushUpExerciseName = 6
	PushUpExerciseNameClappingPushUp                             PushUpExerciseName = 7
	PushUpExerciseNameWeightedClappingPushUp                     PushUpExerciseName = 8
	PushUpExerciseNameCloseGripMedicineBallPushUp                PushUpExerciseName = 9
	PushUpExerciseNameWeightedCloseGripMedicineBallPushUp        PushUpExerciseName = 10
	PushUpExerciseNameCloseHandsPushUp                           PushUpExerciseName = 11
	PushUpExerciseNameWeightedCloseHandsPushUp                   PushUpExerciseName = 12
	PushUpExerciseNameDeclinePushUp                              PushUpExerciseName = 13
	PushUpExerciseNameWeightedDeclinePushUp                      PushUpExerciseName = 14
	PushUpExerciseNameDiamondPushUp                              PushUpExerciseName = 15
	PushUpExerciseNameWeightedDiamondPushUp                      PushUpExerciseName = 16
	PushUpExerciseNameExplosiveCrossoverPushUp                   PushUpExerciseName = 17
	PushUpExerciseNameWeightedExplosiveCrossoverPushUp           PushUpExerciseName = 18
	PushUpExerciseNameExplosivePushUp                            PushUpExerciseName = 19
	PushUpExerciseNameWeightedExplosivePushUp                    PushUpExerciseName = 20
	PushUpExerciseNameFeetElevatedSideToSidePushUp               PushUpExerciseName = 21
	PushUpExerciseNameWeightedFeetElevatedSideToSidePushUp       PushUpExerciseName = 22
	PushUpExerciseNameHandReleasePushUp                          PushUpExerciseName = 23
	PushUpExerciseNameWeightedHandReleasePushUp                  PushUpExerciseName = 24
	PushUpExerciseNameHandstandPushUp                            PushUpExerciseName = 25
	PushUpExerciseNameWeightedHandstandPushUp                    PushUpExerciseName = 26
	PushUpExerciseNameInclinePushUp                              PushUpExerciseName = 27
	PushUpExerciseNameWeightedInclinePushUp                      PushUpExerciseName = 28
	PushUpExerciseNameIsometricExplosivePushUp                   PushUpExerciseName = 29
	PushUpExerciseNameWeightedIsometricExplosivePushUp           PushUpExerciseName = 30
	PushUpExerciseNameJudoPushUp                                 PushUpExerciseName = 31
	PushUpExerciseNameWeightedJudoPushUp                         PushUpExerciseName = 32
	PushUpExerciseNameKneelingPushUp                             PushUpExerciseName = 33
	PushUpExerciseNameWeightedKneelingPushUp                     PushUpExerciseName = 34
	PushUpExerciseNameMedicineBallChestPass                      PushUpExerciseName = 35
	PushUpExerciseNameMedicineBallPushUp                         PushUpExerciseName = 36
	PushUpExerciseNameWeightedMedicineBallPushUp                 PushUpExerciseName = 37
	PushUpExerciseNameOneArmPushUp                               PushUpExerciseName = 38
	PushUpExerciseNameWeightedOneArmPushUp                       PushUpExerciseName = 39
	PushUpExerciseNameWeightedPushUp                             PushUpExerciseName = 40
	PushUpExerciseNamePushUpAndRow                               PushUpExerciseName = 41
	PushUpExerciseNameWeightedPushUpAndRow                       PushUpExerciseName = 42
	PushUpExerciseNamePushUpPlus                                 PushUpExerciseName = 43
	PushUpExerciseNameWeightedPushUpPlus                         PushUpExerciseName = 44
	PushUpExerciseNamePushUpWithFeetOnSwissBall                  PushUpExerciseName = 45
	PushUpExerciseNameWeightedPushUpWithFeetOnSwissBall          PushUpExerciseName = 46
	PushUpExerciseNamePushUpWithOneHandOnMedicineBall            PushUpExerciseName = 47
	PushUpExerciseNameWeightedPushUpWithOneHandOnMedicineBall    PushUpExerciseName = 48
	PushUpExerciseNameShoulderPushUp                             PushUpExerciseName = 49
	PushUpExerciseNameWeightedShoulderPushUp                     PushUpExerciseName = 50
	PushUpExerciseNameSingleArmMedicineBallPushUp                PushUpExerciseName = 51
	PushUpExerciseNameWeightedSingleArmMedicineBallPushUp        PushUpExerciseName = 52
	PushUpExerciseNameSpidermanPushUp                            PushUpExerciseName = 53
	PushUpExerciseNameWeightedSpidermanPushUp                    PushUpExerciseName = 54
	PushUpExerciseNameStackedFeetPushUp                          PushUpExerciseName = 55
	PushUpExerciseNameWeightedStackedFeetPushUp                  PushUpExerciseName = 56
	PushUpExerciseNameStaggeredHandsPushUp                       PushUpExerciseName = 57
	PushUpExerciseNameWeightedStaggeredHandsPushUp               PushUpExerciseName = 58
	PushUpExerciseNameSuspendedPushUp                            PushUpExerciseName = 59
	PushUpExerciseNameWeightedSuspendedPushUp                    PushUpExerciseName = 60
	PushUpExerciseNameSwissBallPushUp                            PushUpExerciseName = 61
	PushUpExerciseNameWeightedSwissBallPushUp                    PushUpExerciseName = 62
	PushUpExerciseNameSwissBallPushUpPlus                        PushUpExerciseName = 63
	PushUpExerciseNameWeightedSwissBallPushUpPlus                PushUpExerciseName = 64
	PushUpExerciseNameTPushUp                                    PushUpExerciseName = 65
	PushUpExerciseNameWeightedTPushUp                            PushUpExerciseName = 66
	PushUpExerciseNameTripleStopPushUp                           PushUpExerciseName = 67
	PushUpExerciseNameWeightedTripleStopPushUp                   PushUpExerciseName = 68
	PushUpExerciseNameWideHandsPushUp                            PushUpExerciseName = 69
	PushUpExerciseNameWeightedWideHandsPushUp                    PushUpExerciseName = 70
	PushUpExerciseNameParalletteHandstandPushUp                  PushUpExerciseName = 71
	PushUpExerciseNameWeightedParalletteHandstandPushUp          PushUpExerciseName = 72
	PushUpExerciseNameRingHandstandPushUp                        PushUpExerciseName = 73
	PushUpExerciseNameWeightedRingHandstandPushUp                PushUpExerciseName = 74
	PushUpExerciseNameRingPushUp                                 PushUpExerciseName = 75
	PushUpExerciseNameWeightedRingPushUp                         PushUpExerciseName = 76
	PushUpExerciseNamePushUp                                     PushUpExerciseName = 77
	PushUpExerciseNamePilatesPushup                              PushUpExerciseName = 78
	PushUpExerciseNameInvalid                                    PushUpExerciseName = 0xFFFF
)

// PwrZoneCalc represents the pwr_zone_calc FIT type.
type PwrZoneCalc byte

const (
	PwrZoneCalcCustom     PwrZoneCalc = 0
	PwrZoneCalcPercentFtp PwrZoneCalc = 1
	PwrZoneCalcInvalid    PwrZoneCalc = 0xFF
)

// RadarThreatLevelType represents the radar_threat_level_type FIT type.
type RadarThreatLevelType byte

const (
	RadarThreatLevelTypeThreatUnknown         RadarThreatLevelType = 0
	RadarThreatLevelTypeThreatNone            RadarThreatLevelType = 1
	RadarThreatLevelTypeThreatApproaching     RadarThreatLevelType = 2
	RadarThreatLevelTypeThreatApproachingFast RadarThreatLevelType = 3
	RadarThreatLevelTypeInvalid               RadarThreatLevelType = 0xFF
)

// RiderPositionType represents the rider_position_type FIT type.
type RiderPositionType byte

const (
	RiderPositionTypeSeated               RiderPositionType = 0
	RiderPositionTypeStanding             RiderPositionType = 1
	RiderPositionTypeTransitionToSeated   RiderPositionType = 2
	RiderPositionTypeTransitionToStanding RiderPositionType = 3
	RiderPositionTypeInvalid              RiderPositionType = 0xFF
)

// RowExerciseName represents the row_exercise_name FIT type.
type RowExerciseName uint16

const (
	RowExerciseNameBarbellStraightLegDeadliftToRow            RowExerciseName = 0
	RowExerciseNameCableRowStanding                           RowExerciseName = 1
	RowExerciseNameDumbbellRow                                RowExerciseName = 2
	RowExerciseNameElevatedFeetInvertedRow                    RowExerciseName = 3
	RowExerciseNameWeightedElevatedFeetInvertedRow            RowExerciseName = 4
	RowExerciseNameFacePull                                   RowExerciseName = 5
	RowExerciseNameFacePullWithExternalRotation               RowExerciseName = 6
	RowExerciseNameInvertedRowWithFeetOnSwissBall             RowExerciseName = 7
	RowExerciseNameWeightedInvertedRowWithFeetOnSwissBall     RowExerciseName = 8
	RowExerciseNameKettlebellRow                              RowExerciseName = 9
	RowExerciseNameModifiedInvertedRow                        RowExerciseName = 10
	RowExerciseNameWeightedModifiedInvertedRow                RowExerciseName = 11
	RowExerciseNameNeutralGripAlternatingDumbbellRow          RowExerciseName = 12
	RowExerciseNameOneArmBentOverRow                          RowExerciseName = 13
	RowExerciseNameOneLeggedDumbbellRow                       RowExerciseName = 14
	RowExerciseNameRenegadeRow                                RowExerciseName = 15
	RowExerciseNameReverseGripBarbellRow                      RowExerciseName = 16
	RowExerciseNameRopeHandleCableRow                         RowExerciseName = 17
	RowExerciseNameSeatedCableRow                             RowExerciseName = 18
	RowExerciseNameSeatedDumbbellRow                          RowExerciseName = 19
	RowExerciseNameSingleArmCableRow                          RowExerciseName = 20
	RowExerciseNameSingleArmCableRowAndRotation               RowExerciseName = 21
	RowExerciseNameSingleArmInvertedRow                       RowExerciseName = 22
	RowExerciseNameWeightedSingleArmInvertedRow               RowExerciseName = 23
	RowExerciseNameSingleArmNeutralGripDumbbellRow            RowExerciseName = 24
	RowExerciseNameSingleArmNeutralGripDumbbellRowAndRotation RowExerciseName = 25
	RowExerciseNameSuspendedInvertedRow                       RowExerciseName = 26
	RowExerciseNameWeightedSuspendedInvertedRow               RowExerciseName = 27
	RowExerciseNameTBarRow                                    RowExerciseName = 28
	RowExerciseNameTowelGripInvertedRow                       RowExerciseName = 29
	RowExerciseNameWeightedTowelGripInvertedRow               RowExerciseName = 30
	RowExerciseNameUnderhandGripCableRow                      RowExerciseName = 31
	RowExerciseNameVGripCableRow                              RowExerciseName = 32
	RowExerciseNameWideGripSeatedCableRow                     RowExerciseName = 33
	RowExerciseNameInvalid                                    RowExerciseName = 0xFFFF
)

// RunExerciseName represents the run_exercise_name FIT type.
type RunExerciseName uint16

const (
	RunExerciseNameRun     RunExerciseName = 0
	RunExerciseNameWalk    RunExerciseName = 1
	RunExerciseNameJog     RunExerciseName = 2
	RunExerciseNameSprint  RunExerciseName = 3
	RunExerciseNameInvalid RunExerciseName = 0xFFFF
)

// Schedule represents the schedule FIT type.
type Schedule byte

const (
	ScheduleWorkout Schedule = 0
	ScheduleCourse  Schedule = 1
	ScheduleInvalid Schedule = 0xFF
)

// SegmentDeleteStatus represents the segment_delete_status FIT type.
type SegmentDeleteStatus byte

const (
	SegmentDeleteStatusDoNotDelete SegmentDeleteStatus = 0
	SegmentDeleteStatusDeleteOne   SegmentDeleteStatus = 1
	SegmentDeleteStatusDeleteAll   SegmentDeleteStatus = 2
	SegmentDeleteStatusInvalid     SegmentDeleteStatus = 0xFF
)

// SegmentLapStatus represents the segment_lap_status FIT type.
type SegmentLapStatus byte

const (
	SegmentLapStatusEnd     SegmentLapStatus = 0
	SegmentLapStatusFail    SegmentLapStatus = 1
	SegmentLapStatusInvalid SegmentLapStatus = 0xFF
)

// SegmentLeaderboardType represents the segment_leaderboard_type FIT type.
type SegmentLeaderboardType byte

const (
	SegmentLeaderboardTypeOverall      SegmentLeaderboardType = 0
	SegmentLeaderboardTypePersonalBest SegmentLeaderboardType = 1
	SegmentLeaderboardTypeConnections  SegmentLeaderboardType = 2
	SegmentLeaderboardTypeGroup        SegmentLeaderboardType = 3
	SegmentLeaderboardTypeChallenger   SegmentLeaderboardType = 4
	SegmentLeaderboardTypeKom          SegmentLeaderboardType = 5
	SegmentLeaderboardTypeQom          SegmentLeaderboardType = 6
	SegmentLeaderboardTypePr           SegmentLeaderboardType = 7
	SegmentLeaderboardTypeGoal         SegmentLeaderboardType = 8
	SegmentLeaderboardTypeRival        SegmentLeaderboardType = 9
	SegmentLeaderboardTypeClubLeader   SegmentLeaderboardType = 10
	SegmentLeaderboardTypeInvalid      SegmentLeaderboardType = 0xFF
)

// SegmentSelectionType represents the segment_selection_type FIT type.
type SegmentSelectionType byte

const (
	SegmentSelectionTypeStarred   SegmentSelectionType = 0
	SegmentSelectionTypeSuggested SegmentSelectionType = 1
	SegmentSelectionTypeInvalid   SegmentSelectionType = 0xFF
)

// SensorType represents the sensor_type FIT type.
type SensorType byte

const (
	SensorTypeAccelerometer SensorType = 0
	SensorTypeGyroscope     SensorType = 1
	SensorTypeCompass       SensorType = 2 // Magnetometer
	SensorTypeBarometer     SensorType = 3
	SensorTypeInvalid       SensorType = 0xFF
)

// SessionTrigger represents the session_trigger FIT type.
type SessionTrigger byte

const (
	SessionTriggerActivityEnd      SessionTrigger = 0
	SessionTriggerManual           SessionTrigger = 1 // User changed sport.
	SessionTriggerAutoMultiSport   SessionTrigger = 2 // Auto multi-sport feature is enabled and user pressed lap button to advance session.
	SessionTriggerFitnessEquipment SessionTrigger = 3 // Auto sport change caused by user linking to fitness equipment.
	SessionTriggerInvalid          SessionTrigger = 0xFF
)

// SetType represents the set_type FIT type.
type SetType uint8

const (
	SetTypeRest    SetType = 0
	SetTypeActive  SetType = 1
	SetTypeInvalid SetType = 0xFF
)

// ShoulderPressExerciseName represents the shoulder_press_exercise_name FIT type.
type ShoulderPressExerciseName uint16

const (
	ShoulderPressExerciseNameAlternatingDumbbellShoulderPress         ShoulderPressExerciseName = 0
	ShoulderPressExerciseNameArnoldPress                              ShoulderPressExerciseName = 1
	ShoulderPressExerciseNameBarbellFrontSquatToPushPress             ShoulderPressExerciseName = 2
	ShoulderPressExerciseNameBarbellPushPress                         ShoulderPressExerciseName = 3
	ShoulderPressExerciseNameBarbellShoulderPress                     ShoulderPressExerciseName = 4
	ShoulderPressExerciseNameDeadCurlPress                            ShoulderPressExerciseName = 5
	ShoulderPressExerciseNameDumbbellAlternatingShoulderPressAndTwist ShoulderPressExerciseName = 6
	ShoulderPressExerciseNameDumbbellHammerCurlToLungeToPress         ShoulderPressExerciseName = 7
	ShoulderPressExerciseNameDumbbellPushPress                        ShoulderPressExerciseName = 8
	ShoulderPressExerciseNameFloorInvertedShoulderPress               ShoulderPressExerciseName = 9
	ShoulderPressExerciseNameWeightedFloorInvertedShoulderPress       ShoulderPressExerciseName = 10
	ShoulderPressExerciseNameInvertedShoulderPress                    ShoulderPressExerciseName = 11
	ShoulderPressExerciseNameWeightedInvertedShoulderPress            ShoulderPressExerciseName = 12
	ShoulderPressExerciseNameOneArmPushPress                          ShoulderPressExerciseName = 13
	ShoulderPressExerciseNameOverheadBarbellPress                     ShoulderPressExerciseName = 14
	ShoulderPressExerciseNameOverheadDumbbellPress                    ShoulderPressExerciseName = 15
	ShoulderPressExerciseNameSeatedBarbellShoulderPress               ShoulderPressExerciseName = 16
	ShoulderPressExerciseNameSeatedDumbbellShoulderPress              ShoulderPressExerciseName = 17
	ShoulderPressExerciseNameSingleArmDumbbellShoulderPress           ShoulderPressExerciseName = 18
	ShoulderPressExerciseNameSingleArmStepUpAndPress                  ShoulderPressExerciseName = 19
	ShoulderPressExerciseNameSmithMachineOverheadPress                ShoulderPressExerciseName = 20
	ShoulderPressExerciseNameSplitStanceHammerCurlToPress             ShoulderPressExerciseName = 21
	ShoulderPressExerciseNameSwissBallDumbbellShoulderPress           ShoulderPressExerciseName = 22
	ShoulderPressExerciseNameWeightPlateFrontRaise                    ShoulderPressExerciseName = 23
	ShoulderPressExerciseNameInvalid                                  ShoulderPressExerciseName = 0xFFFF
)

// ShoulderStabilityExerciseName represents the shoulder_stability_exercise_name FIT type.
type ShoulderStabilityExerciseName uint16

const (
	ShoulderStabilityExerciseName90DegreeCableExternalRotation          ShoulderStabilityExerciseName = 0
	ShoulderStabilityExerciseNameBandExternalRotation                   ShoulderStabilityExerciseName = 1
	ShoulderStabilityExerciseNameBandInternalRotation                   ShoulderStabilityExerciseName = 2
	ShoulderStabilityExerciseNameBentArmLateralRaiseAndExternalRotation ShoulderStabilityExerciseName = 3
	ShoulderStabilityExerciseNameCableExternalRotation                  ShoulderStabilityExerciseName = 4
	ShoulderStabilityExerciseNameDumbbellFacePullWithExternalRotation   ShoulderStabilityExerciseName = 5
	ShoulderStabilityExerciseNameFloorIRaise                            ShoulderStabilityExerciseName = 6
	ShoulderStabilityExerciseNameWeightedFloorIRaise                    ShoulderStabilityExerciseName = 7
	ShoulderStabilityExerciseNameFloorTRaise                            ShoulderStabilityExerciseName = 8
	ShoulderStabilityExerciseNameWeightedFloorTRaise                    ShoulderStabilityExerciseName = 9
	ShoulderStabilityExerciseNameFloorYRaise                            ShoulderStabilityExerciseName = 10
	ShoulderStabilityExerciseNameWeightedFloorYRaise                    ShoulderStabilityExerciseName = 11
	ShoulderStabilityExerciseNameInclineIRaise                          ShoulderStabilityExerciseName = 12
	ShoulderStabilityExerciseNameWeightedInclineIRaise                  ShoulderStabilityExerciseName = 13
	ShoulderStabilityExerciseNameInclineLRaise                          ShoulderStabilityExerciseName = 14
	ShoulderStabilityExerciseNameWeightedInclineLRaise                  ShoulderStabilityExerciseName = 15
	ShoulderStabilityExerciseNameInclineTRaise                          ShoulderStabilityExerciseName = 16
	ShoulderStabilityExerciseNameWeightedInclineTRaise                  ShoulderStabilityExerciseName = 17
	ShoulderStabilityExerciseNameInclineWRaise                          ShoulderStabilityExerciseName = 18
	ShoulderStabilityExerciseNameWeightedInclineWRaise                  ShoulderStabilityExerciseName = 19
	ShoulderStabilityExerciseNameInclineYRaise                          ShoulderStabilityExerciseName = 20
	ShoulderStabilityExerciseNameWeightedInclineYRaise                  ShoulderStabilityExerciseName = 21
	ShoulderStabilityExerciseNameLyingExternalRotation                  ShoulderStabilityExerciseName = 22
	ShoulderStabilityExerciseNameSeatedDumbbellExternalRotation         ShoulderStabilityExerciseName = 23
	ShoulderStabilityExerciseNameStandingLRaise                         ShoulderStabilityExerciseName = 24
	ShoulderStabilityExerciseNameSwissBallIRaise                        ShoulderStabilityExerciseName = 25
	ShoulderStabilityExerciseNameWeightedSwissBallIRaise                ShoulderStabilityExerciseName = 26
	ShoulderStabilityExerciseNameSwissBallTRaise                        ShoulderStabilityExerciseName = 27
	ShoulderStabilityExerciseNameWeightedSwissBallTRaise                ShoulderStabilityExerciseName = 28
	ShoulderStabilityExerciseNameSwissBallWRaise                        ShoulderStabilityExerciseName = 29
	ShoulderStabilityExerciseNameWeightedSwissBallWRaise                ShoulderStabilityExerciseName = 30
	ShoulderStabilityExerciseNameSwissBallYRaise                        ShoulderStabilityExerciseName = 31
	ShoulderStabilityExerciseNameWeightedSwissBallYRaise                ShoulderStabilityExerciseName = 32
	ShoulderStabilityExerciseNameInvalid                                ShoulderStabilityExerciseName = 0xFFFF
)

// ShrugExerciseName represents the shrug_exercise_name FIT type.
type ShrugExerciseName uint16

const (
	ShrugExerciseNameBarbellJumpShrug               ShrugExerciseName = 0
	ShrugExerciseNameBarbellShrug                   ShrugExerciseName = 1
	ShrugExerciseNameBarbellUprightRow              ShrugExerciseName = 2
	ShrugExerciseNameBehindTheBackSmithMachineShrug ShrugExerciseName = 3
	ShrugExerciseNameDumbbellJumpShrug              ShrugExerciseName = 4
	ShrugExerciseNameDumbbellShrug                  ShrugExerciseName = 5
	ShrugExerciseNameDumbbellUprightRow             ShrugExerciseName = 6
	ShrugExerciseNameInclineDumbbellShrug           ShrugExerciseName = 7
	ShrugExerciseNameOverheadBarbellShrug           ShrugExerciseName = 8
	ShrugExerciseNameOverheadDumbbellShrug          ShrugExerciseName = 9
	ShrugExerciseNameScaptionAndShrug               ShrugExerciseName = 10
	ShrugExerciseNameScapularRetraction             ShrugExerciseName = 11
	ShrugExerciseNameSerratusChairShrug             ShrugExerciseName = 12
	ShrugExerciseNameWeightedSerratusChairShrug     ShrugExerciseName = 13
	ShrugExerciseNameSerratusShrug                  ShrugExerciseName = 14
	ShrugExerciseNameWeightedSerratusShrug          ShrugExerciseName = 15
	ShrugExerciseNameWideGripJumpShrug              ShrugExerciseName = 16
	ShrugExerciseNameInvalid                        ShrugExerciseName = 0xFFFF
)

// Side represents the side FIT type.
type Side byte

const (
	SideRight   Side = 0
	SideLeft    Side = 1
	SideInvalid Side = 0xFF
)

// SitUpExerciseName represents the sit_up_exercise_name FIT type.
type SitUpExerciseName uint16

const (
	SitUpExerciseNameAlternatingSitUp                    SitUpExerciseName = 0
	SitUpExerciseNameWeightedAlternatingSitUp            SitUpExerciseName = 1
	SitUpExerciseNameBentKneeVUp                         SitUpExerciseName = 2
	SitUpExerciseNameWeightedBentKneeVUp                 SitUpExerciseName = 3
	SitUpExerciseNameButterflySitUp                      SitUpExerciseName = 4
	SitUpExerciseNameWeightedButterflySitup              SitUpExerciseName = 5
	SitUpExerciseNameCrossPunchRollUp                    SitUpExerciseName = 6
	SitUpExerciseNameWeightedCrossPunchRollUp            SitUpExerciseName = 7
	SitUpExerciseNameCrossedArmsSitUp                    SitUpExerciseName = 8
	SitUpExerciseNameWeightedCrossedArmsSitUp            SitUpExerciseName = 9
	SitUpExerciseNameGetUpSitUp                          SitUpExerciseName = 10
	SitUpExerciseNameWeightedGetUpSitUp                  SitUpExerciseName = 11
	SitUpExerciseNameHoveringSitUp                       SitUpExerciseName = 12
	SitUpExerciseNameWeightedHoveringSitUp               SitUpExerciseName = 13
	SitUpExerciseNameKettlebellSitUp                     SitUpExerciseName = 14
	SitUpExerciseNameMedicineBallAlternatingVUp          SitUpExerciseName = 15
	SitUpExerciseNameMedicineBallSitUp                   SitUpExerciseName = 16
	SitUpExerciseNameMedicineBallVUp                     SitUpExerciseName = 17
	SitUpExerciseNameModifiedSitUp                       SitUpExerciseName = 18
	SitUpExerciseNameNegativeSitUp                       SitUpExerciseName = 19
	SitUpExerciseNameOneArmFullSitUp                     SitUpExerciseName = 20
	SitUpExerciseNameRecliningCircle                     SitUpExerciseName = 21
	SitUpExerciseNameWeightedRecliningCircle             SitUpExerciseName = 22
	SitUpExerciseNameReverseCurlUp                       SitUpExerciseName = 23
	SitUpExerciseNameWeightedReverseCurlUp               SitUpExerciseName = 24
	SitUpExerciseNameSingleLegSwissBallJackknife         SitUpExerciseName = 25
	SitUpExerciseNameWeightedSingleLegSwissBallJackknife SitUpExerciseName = 26
	SitUpExerciseNameTheTeaser                           SitUpExerciseName = 27
	SitUpExerciseNameTheTeaserWeighted                   SitUpExerciseName = 28
	SitUpExerciseNameThreePartRollDown                   SitUpExerciseName = 29
	SitUpExerciseNameWeightedThreePartRollDown           SitUpExerciseName = 30
	SitUpExerciseNameVUp                                 SitUpExerciseName = 31
	SitUpExerciseNameWeightedVUp                         SitUpExerciseName = 32
	SitUpExerciseNameWeightedRussianTwistOnSwissBall     SitUpExerciseName = 33
	SitUpExerciseNameWeightedSitUp                       SitUpExerciseName = 34
	SitUpExerciseNameXAbs                                SitUpExerciseName = 35
	SitUpExerciseNameWeightedXAbs                        SitUpExerciseName = 36
	SitUpExerciseNameSitUp                               SitUpExerciseName = 37
	SitUpExerciseNameInvalid                             SitUpExerciseName = 0xFFFF
)

// SleepLevel represents the sleep_level FIT type.
type SleepLevel byte

const (
	SleepLevelUnmeasurable SleepLevel = 0
	SleepLevelAwake        SleepLevel = 1
	SleepLevelLight        SleepLevel = 2
	SleepLevelDeep         SleepLevel = 3
	SleepLevelRem          SleepLevel = 4
	SleepLevelInvalid      SleepLevel = 0xFF
)

// SourceType represents the source_type FIT type.
type SourceType byte

const (
	SourceTypeAnt                SourceType = 0 // External device connected with ANT
	SourceTypeAntplus            SourceType = 1 // External device connected with ANT+
	SourceTypeBluetooth          SourceType = 2 // External device connected with BT
	SourceTypeBluetoothLowEnergy SourceType = 3 // External device connected with BLE
	SourceTypeWifi               SourceType = 4 // External device connected with Wifi
	SourceTypeLocal              SourceType = 5 // Onboard device
	SourceTypeInvalid            SourceType = 0xFF
)

// SplitType represents the split_type FIT type.
type SplitType byte

const (
	SplitTypeAscentSplit      SplitType = 1
	SplitTypeDescentSplit     SplitType = 2
	SplitTypeIntervalActive   SplitType = 3
	SplitTypeIntervalRest     SplitType = 4
	SplitTypeIntervalWarmup   SplitType = 5
	SplitTypeIntervalCooldown SplitType = 6
	SplitTypeIntervalRecovery SplitType = 7
	SplitTypeIntervalOther    SplitType = 8
	SplitTypeClimbActive      SplitType = 9
	SplitTypeClimbRest        SplitType = 10
	SplitTypeSurfActive       SplitType = 11
	SplitTypeRunActive        SplitType = 12
	SplitTypeRunRest          SplitType = 13
	SplitTypeWorkoutRound     SplitType = 14
	SplitTypeRwdRun           SplitType = 17 // run/walk detection running
	SplitTypeRwdWalk          SplitType = 18 // run/walk detection walking
	SplitTypeWindsurfActive   SplitType = 21
	SplitTypeRwdStand         SplitType = 22 // run/walk detection standing
	SplitTypeTransition       SplitType = 23 // Marks the time going from ascent_split to descent_split/used in backcountry ski
	SplitTypeSkiLiftSplit     SplitType = 28
	SplitTypeSkiRunSplit      SplitType = 29
	SplitTypeInvalid          SplitType = 0xFF
)

// Spo2MeasurementType represents the spo2_measurement_type FIT type.
type Spo2MeasurementType byte

const (
	Spo2MeasurementTypeOffWrist        Spo2MeasurementType = 0
	Spo2MeasurementTypeSpotCheck       Spo2MeasurementType = 1
	Spo2MeasurementTypeContinuousCheck Spo2MeasurementType = 2
	Spo2MeasurementTypePeriodic        Spo2MeasurementType = 3
	Spo2MeasurementTypeInvalid         Spo2MeasurementType = 0xFF
)

// Sport represents the sport FIT type.
type Sport byte

const (
	SportGeneric               Sport = 0
	SportRunning               Sport = 1
	SportCycling               Sport = 2
	SportTransition            Sport = 3 // Mulitsport transition
	SportFitnessEquipment      Sport = 4
	SportSwimming              Sport = 5
	SportBasketball            Sport = 6
	SportSoccer                Sport = 7
	SportTennis                Sport = 8
	SportAmericanFootball      Sport = 9
	SportTraining              Sport = 10
	SportWalking               Sport = 11
	SportCrossCountrySkiing    Sport = 12
	SportAlpineSkiing          Sport = 13
	SportSnowboarding          Sport = 14
	SportRowing                Sport = 15
	SportMountaineering        Sport = 16
	SportHiking                Sport = 17
	SportMultisport            Sport = 18
	SportPaddling              Sport = 19
	SportFlying                Sport = 20
	SportEBiking               Sport = 21
	SportMotorcycling          Sport = 22
	SportBoating               Sport = 23
	SportDriving               Sport = 24
	SportGolf                  Sport = 25
	SportHangGliding           Sport = 26
	SportHorsebackRiding       Sport = 27
	SportHunting               Sport = 28
	SportFishing               Sport = 29
	SportInlineSkating         Sport = 30
	SportRockClimbing          Sport = 31
	SportSailing               Sport = 32
	SportIceSkating            Sport = 33
	SportSkyDiving             Sport = 34
	SportSnowshoeing           Sport = 35
	SportSnowmobiling          Sport = 36
	SportStandUpPaddleboarding Sport = 37
	SportSurfing               Sport = 38
	SportWakeboarding          Sport = 39
	SportWaterSkiing           Sport = 40
	SportKayaking              Sport = 41
	SportRafting               Sport = 42
	SportWindsurfing           Sport = 43
	SportKitesurfing           Sport = 44
	SportTactical              Sport = 45
	SportJumpmaster            Sport = 46
	SportBoxing                Sport = 47
	SportFloorClimbing         Sport = 48
	SportDiving                Sport = 53
	SportHiit                  Sport = 62
	SportRacket                Sport = 64
	SportWaterTubing           Sport = 76
	SportWakesurfing           Sport = 77
	SportAll                   Sport = 254 // All is for goals only to include all sports.
	SportInvalid               Sport = 0xFF
)

// SportBits0 represents the sport_bits_0 FIT type.
type SportBits0 uint8

const (
	SportBits0Generic          SportBits0 = 0x01
	SportBits0Running          SportBits0 = 0x02
	SportBits0Cycling          SportBits0 = 0x04
	SportBits0Transition       SportBits0 = 0x08 // Mulitsport transition
	SportBits0FitnessEquipment SportBits0 = 0x10
	SportBits0Swimming         SportBits0 = 0x20
	SportBits0Basketball       SportBits0 = 0x40
	SportBits0Soccer           SportBits0 = 0x80
	SportBits0Invalid          SportBits0 = 0x00
)

// SportBits1 represents the sport_bits_1 FIT type.
type SportBits1 uint8

const (
	SportBits1Tennis             SportBits1 = 0x01
	SportBits1AmericanFootball   SportBits1 = 0x02
	SportBits1Training           SportBits1 = 0x04
	SportBits1Walking            SportBits1 = 0x08
	SportBits1CrossCountrySkiing SportBits1 = 0x10
	SportBits1AlpineSkiing       SportBits1 = 0x20
	SportBits1Snowboarding       SportBits1 = 0x40
	SportBits1Rowing             SportBits1 = 0x80
	SportBits1Invalid            SportBits1 = 0x00
)

// SportBits2 represents the sport_bits_2 FIT type.
type SportBits2 uint8

const (
	SportBits2Mountaineering SportBits2 = 0x01
	SportBits2Hiking         SportBits2 = 0x02
	SportBits2Multisport     SportBits2 = 0x04
	SportBits2Paddling       SportBits2 = 0x08
	SportBits2Flying         SportBits2 = 0x10
	SportBits2EBiking        SportBits2 = 0x20
	SportBits2Motorcycling   SportBits2 = 0x40
	SportBits2Boating        SportBits2 = 0x80
	SportBits2Invalid        SportBits2 = 0x00
)

// SportBits3 represents the sport_bits_3 FIT type.
type SportBits3 uint8

const (
	SportBits3Driving         SportBits3 = 0x01
	SportBits3Golf            SportBits3 = 0x02
	SportBits3HangGliding     SportBits3 = 0x04
	SportBits3HorsebackRiding SportBits3 = 0x08
	SportBits3Hunting         SportBits3 = 0x10
	SportBits3Fishing         SportBits3 = 0x20
	SportBits3InlineSkating   SportBits3 = 0x40
	SportBits3RockClimbing    SportBits3 = 0x80
	SportBits3Invalid         SportBits3 = 0x00
)

// SportBits4 represents the sport_bits_4 FIT type.
type SportBits4 uint8

const (
	SportBits4Sailing               SportBits4 = 0x01
	SportBits4IceSkating            SportBits4 = 0x02
	SportBits4SkyDiving             SportBits4 = 0x04
	SportBits4Snowshoeing           SportBits4 = 0x08
	SportBits4Snowmobiling          SportBits4 = 0x10
	SportBits4StandUpPaddleboarding SportBits4 = 0x20
	SportBits4Surfing               SportBits4 = 0x40
	SportBits4Wakeboarding          SportBits4 = 0x80
	SportBits4Invalid               SportBits4 = 0x00
)

// SportBits5 represents the sport_bits_5 FIT type.
type SportBits5 uint8

const (
	SportBits5WaterSkiing SportBits5 = 0x01
	SportBits5Kayaking    SportBits5 = 0x02
	SportBits5Rafting     SportBits5 = 0x04
	SportBits5Windsurfing SportBits5 = 0x08
	SportBits5Kitesurfing SportBits5 = 0x10
	SportBits5Tactical    SportBits5 = 0x20
	SportBits5Jumpmaster  SportBits5 = 0x40
	SportBits5Boxing      SportBits5 = 0x80
	SportBits5Invalid     SportBits5 = 0x00
)

// SportBits6 represents the sport_bits_6 FIT type.
type SportBits6 uint8

const (
	SportBits6FloorClimbing SportBits6 = 0x01
	SportBits6Invalid       SportBits6 = 0x00
)

// SportEvent represents the sport_event FIT type.
type SportEvent byte

const (
	SportEventUncategorized  SportEvent = 0
	SportEventGeocaching     SportEvent = 1
	SportEventFitness        SportEvent = 2
	SportEventRecreation     SportEvent = 3
	SportEventRace           SportEvent = 4
	SportEventSpecialEvent   SportEvent = 5
	SportEventTraining       SportEvent = 6
	SportEventTransportation SportEvent = 7
	SportEventTouring        SportEvent = 8
	SportEventInvalid        SportEvent = 0xFF
)

// SquatExerciseName represents the squat_exercise_name FIT type.
type SquatExerciseName uint16

const (
	SquatExerciseNameLegPress                                        SquatExerciseName = 0
	SquatExerciseNameBackSquatWithBodyBar                            SquatExerciseName = 1
	SquatExerciseNameBackSquats                                      SquatExerciseName = 2
	SquatExerciseNameWeightedBackSquats                              SquatExerciseName = 3
	SquatExerciseNameBalancingSquat                                  SquatExerciseName = 4
	SquatExerciseNameWeightedBalancingSquat                          SquatExerciseName = 5
	SquatExerciseNameBarbellBackSquat                                SquatExerciseName = 6
	SquatExerciseNameBarbellBoxSquat                                 SquatExerciseName = 7
	SquatExerciseNameBarbellFrontSquat                               SquatExerciseName = 8
	SquatExerciseNameBarbellHackSquat                                SquatExerciseName = 9
	SquatExerciseNameBarbellHangSquatSnatch                          SquatExerciseName = 10
	SquatExerciseNameBarbellLateralStepUp                            SquatExerciseName = 11
	SquatExerciseNameBarbellQuarterSquat                             SquatExerciseName = 12
	SquatExerciseNameBarbellSiffSquat                                SquatExerciseName = 13
	SquatExerciseNameBarbellSquatSnatch                              SquatExerciseName = 14
	SquatExerciseNameBarbellSquatWithHeelsRaised                     SquatExerciseName = 15
	SquatExerciseNameBarbellStepover                                 SquatExerciseName = 16
	SquatExerciseNameBarbellStepUp                                   SquatExerciseName = 17
	SquatExerciseNameBenchSquatWithRotationalChop                    SquatExerciseName = 18
	SquatExerciseNameWeightedBenchSquatWithRotationalChop            SquatExerciseName = 19
	SquatExerciseNameBodyWeightWallSquat                             SquatExerciseName = 20
	SquatExerciseNameWeightedWallSquat                               SquatExerciseName = 21
	SquatExerciseNameBoxStepSquat                                    SquatExerciseName = 22
	SquatExerciseNameWeightedBoxStepSquat                            SquatExerciseName = 23
	SquatExerciseNameBracedSquat                                     SquatExerciseName = 24
	SquatExerciseNameCrossedArmBarbellFrontSquat                     SquatExerciseName = 25
	SquatExerciseNameCrossoverDumbbellStepUp                         SquatExerciseName = 26
	SquatExerciseNameDumbbellFrontSquat                              SquatExerciseName = 27
	SquatExerciseNameDumbbellSplitSquat                              SquatExerciseName = 28
	SquatExerciseNameDumbbellSquat                                   SquatExerciseName = 29
	SquatExerciseNameDumbbellSquatClean                              SquatExerciseName = 30
	SquatExerciseNameDumbbellStepover                                SquatExerciseName = 31
	SquatExerciseNameDumbbellStepUp                                  SquatExerciseName = 32
	SquatExerciseNameElevatedSingleLegSquat                          SquatExerciseName = 33
	SquatExerciseNameWeightedElevatedSingleLegSquat                  SquatExerciseName = 34
	SquatExerciseNameFigureFourSquats                                SquatExerciseName = 35
	SquatExerciseNameWeightedFigureFourSquats                        SquatExerciseName = 36
	SquatExerciseNameGobletSquat                                     SquatExerciseName = 37
	SquatExerciseNameKettlebellSquat                                 SquatExerciseName = 38
	SquatExerciseNameKettlebellSwingOverhead                         SquatExerciseName = 39
	SquatExerciseNameKettlebellSwingWithFlipToSquat                  SquatExerciseName = 40
	SquatExerciseNameLateralDumbbellStepUp                           SquatExerciseName = 41
	SquatExerciseNameOneLeggedSquat                                  SquatExerciseName = 42
	SquatExerciseNameOverheadDumbbellSquat                           SquatExerciseName = 43
	SquatExerciseNameOverheadSquat                                   SquatExerciseName = 44
	SquatExerciseNamePartialSingleLegSquat                           SquatExerciseName = 45
	SquatExerciseNameWeightedPartialSingleLegSquat                   SquatExerciseName = 46
	SquatExerciseNamePistolSquat                                     SquatExerciseName = 47
	SquatExerciseNameWeightedPistolSquat                             SquatExerciseName = 48
	SquatExerciseNamePlieSlides                                      SquatExerciseName = 49
	SquatExerciseNameWeightedPlieSlides                              SquatExerciseName = 50
	SquatExerciseNamePlieSquat                                       SquatExerciseName = 51
	SquatExerciseNameWeightedPlieSquat                               SquatExerciseName = 52
	SquatExerciseNamePrisonerSquat                                   SquatExerciseName = 53
	SquatExerciseNameWeightedPrisonerSquat                           SquatExerciseName = 54
	SquatExerciseNameSingleLegBenchGetUp                             SquatExerciseName = 55
	SquatExerciseNameWeightedSingleLegBenchGetUp                     SquatExerciseName = 56
	SquatExerciseNameSingleLegBenchSquat                             SquatExerciseName = 57
	SquatExerciseNameWeightedSingleLegBenchSquat                     SquatExerciseName = 58
	SquatExerciseNameSingleLegSquatOnSwissBall                       SquatExerciseName = 59
	SquatExerciseNameWeightedSingleLegSquatOnSwissBall               SquatExerciseName = 60
	SquatExerciseNameSquat                                           SquatExerciseName = 61
	SquatExerciseNameWeightedSquat                                   SquatExerciseName = 62
	SquatExerciseNameSquatsWithBand                                  SquatExerciseName = 63
	SquatExerciseNameStaggeredSquat                                  SquatExerciseName = 64
	SquatExerciseNameWeightedStaggeredSquat                          SquatExerciseName = 65
	SquatExerciseNameStepUp                                          SquatExerciseName = 66
	SquatExerciseNameWeightedStepUp                                  SquatExerciseName = 67
	SquatExerciseNameSuitcaseSquats                                  SquatExerciseName = 68
	SquatExerciseNameSumoSquat                                       SquatExerciseName = 69
	SquatExerciseNameSumoSquatSlideIn                                SquatExerciseName = 70
	SquatExerciseNameWeightedSumoSquatSlideIn                        SquatExerciseName = 71
	SquatExerciseNameSumoSquatToHighPull                             SquatExerciseName = 72
	SquatExerciseNameSumoSquatToStand                                SquatExerciseName = 73
	SquatExerciseNameWeightedSumoSquatToStand                        SquatExerciseName = 74
	SquatExerciseNameSumoSquatWithRotation                           SquatExerciseName = 75
	SquatExerciseNameWeightedSumoSquatWithRotation                   SquatExerciseName = 76
	SquatExerciseNameSwissBallBodyWeightWallSquat                    SquatExerciseName = 77
	SquatExerciseNameWeightedSwissBallWallSquat                      SquatExerciseName = 78
	SquatExerciseNameThrusters                                       SquatExerciseName = 79
	SquatExerciseNameUnevenSquat                                     SquatExerciseName = 80
	SquatExerciseNameWeightedUnevenSquat                             SquatExerciseName = 81
	SquatExerciseNameWaistSlimmingSquat                              SquatExerciseName = 82
	SquatExerciseNameWallBall                                        SquatExerciseName = 83
	SquatExerciseNameWideStanceBarbellSquat                          SquatExerciseName = 84
	SquatExerciseNameWideStanceGobletSquat                           SquatExerciseName = 85
	SquatExerciseNameZercherSquat                                    SquatExerciseName = 86
	SquatExerciseNameKbsOverhead                                     SquatExerciseName = 87 // Deprecated do not use
	SquatExerciseNameSquatAndSideKick                                SquatExerciseName = 88
	SquatExerciseNameSquatJumpsInNOut                                SquatExerciseName = 89
	SquatExerciseNamePilatesPlieSquatsParallelTurnedOutFlatAndHeels  SquatExerciseName = 90
	SquatExerciseNameReleveStraightLegAndKneeBentWithOneLegVariation SquatExerciseName = 91
	SquatExerciseNameInvalid                                         SquatExerciseName = 0xFFFF
)

// StrokeType represents the stroke_type FIT type.
type StrokeType byte

const (
	StrokeTypeNoEvent  StrokeType = 0
	StrokeTypeOther    StrokeType = 1 // stroke was detected but cannot be identified
	StrokeTypeServe    StrokeType = 2
	StrokeTypeForehand StrokeType = 3
	StrokeTypeBackhand StrokeType = 4
	StrokeTypeSmash    StrokeType = 5
	StrokeTypeInvalid  StrokeType = 0xFF
)

// SubSport represents the sub_sport FIT type.
type SubSport byte

const (
	SubSportGeneric              SubSport = 0
	SubSportTreadmill            SubSport = 1  // Run/Fitness Equipment
	SubSportStreet               SubSport = 2  // Run
	SubSportTrail                SubSport = 3  // Run
	SubSportTrack                SubSport = 4  // Run
	SubSportSpin                 SubSport = 5  // Cycling
	SubSportIndoorCycling        SubSport = 6  // Cycling/Fitness Equipment
	SubSportRoad                 SubSport = 7  // Cycling
	SubSportMountain             SubSport = 8  // Cycling
	SubSportDownhill             SubSport = 9  // Cycling
	SubSportRecumbent            SubSport = 10 // Cycling
	SubSportCyclocross           SubSport = 11 // Cycling
	SubSportHandCycling          SubSport = 12 // Cycling
	SubSportTrackCycling         SubSport = 13 // Cycling
	SubSportIndoorRowing         SubSport = 14 // Fitness Equipment
	SubSportElliptical           SubSport = 15 // Fitness Equipment
	SubSportStairClimbing        SubSport = 16 // Fitness Equipment
	SubSportLapSwimming          SubSport = 17 // Swimming
	SubSportOpenWater            SubSport = 18 // Swimming
	SubSportFlexibilityTraining  SubSport = 19 // Training
	SubSportStrengthTraining     SubSport = 20 // Training
	SubSportWarmUp               SubSport = 21 // Tennis
	SubSportMatch                SubSport = 22 // Tennis
	SubSportExercise             SubSport = 23 // Tennis
	SubSportChallenge            SubSport = 24
	SubSportIndoorSkiing         SubSport = 25 // Fitness Equipment
	SubSportCardioTraining       SubSport = 26 // Training
	SubSportIndoorWalking        SubSport = 27 // Walking/Fitness Equipment
	SubSportEBikeFitness         SubSport = 28 // E-Biking
	SubSportBmx                  SubSport = 29 // Cycling
	SubSportCasualWalking        SubSport = 30 // Walking
	SubSportSpeedWalking         SubSport = 31 // Walking
	SubSportBikeToRunTransition  SubSport = 32 // Transition
	SubSportRunToBikeTransition  SubSport = 33 // Transition
	SubSportSwimToBikeTransition SubSport = 34 // Transition
	SubSportAtv                  SubSport = 35 // Motorcycling
	SubSportMotocross            SubSport = 36 // Motorcycling
	SubSportBackcountry          SubSport = 37 // Alpine Skiing/Snowboarding
	SubSportResort               SubSport = 38 // Alpine Skiing/Snowboarding
	SubSportRcDrone              SubSport = 39 // Flying
	SubSportWingsuit             SubSport = 40 // Flying
	SubSportWhitewater           SubSport = 41 // Kayaking/Rafting
	SubSportSkateSkiing          SubSport = 42 // Cross Country Skiing
	SubSportYoga                 SubSport = 43 // Training
	SubSportPilates              SubSport = 44 // Fitness Equipment
	SubSportIndoorRunning        SubSport = 45 // Run
	SubSportGravelCycling        SubSport = 46 // Cycling
	SubSportEBikeMountain        SubSport = 47 // Cycling
	SubSportCommuting            SubSport = 48 // Cycling
	SubSportMixedSurface         SubSport = 49 // Cycling
	SubSportNavigate             SubSport = 50
	SubSportTrackMe              SubSport = 51
	SubSportMap                  SubSport = 52
	SubSportSingleGasDiving      SubSport = 53 // Diving
	SubSportMultiGasDiving       SubSport = 54 // Diving
	SubSportGaugeDiving          SubSport = 55 // Diving
	SubSportApneaDiving          SubSport = 56 // Diving
	SubSportApneaHunting         SubSport = 57 // Diving
	SubSportVirtualActivity      SubSport = 58
	SubSportObstacle             SubSport = 59 // Used for events where participants run, crawl through mud, climb over walls, etc.
	SubSportBreathing            SubSport = 62
	SubSportSailRace             SubSport = 65  // Sailing
	SubSportUltra                SubSport = 67  // Ultramarathon
	SubSportIndoorClimbing       SubSport = 68  // Climbing
	SubSportBouldering           SubSport = 69  // Climbing
	SubSportHiit                 SubSport = 70  // High Intensity Interval Training
	SubSportAmrap                SubSport = 73  // HIIT
	SubSportEmom                 SubSport = 74  // HIIT
	SubSportTabata               SubSport = 75  // HIIT
	SubSportPickleball           SubSport = 84  // Racket
	SubSportPadel                SubSport = 85  // Racket
	SubSportFlyCanopy            SubSport = 110 // Flying
	SubSportFlyParaglide         SubSport = 111 // Flying
	SubSportFlyParamotor         SubSport = 112 // Flying
	SubSportFlyPressurized       SubSport = 113 // Flying
	SubSportFlyNavigate          SubSport = 114 // Flying
	SubSportFlyTimer             SubSport = 115 // Flying
	SubSportFlyAltimeter         SubSport = 116 // Flying
	SubSportFlyWx                SubSport = 117 // Flying
	SubSportFlyVfr               SubSport = 118 // Flying
	SubSportFlyIfr               SubSport = 119 // Flying
	SubSportAll                  SubSport = 254
	SubSportInvalid              SubSport = 0xFF
)

// SupportedExdScreenLayouts represents the supported_exd_screen_layouts FIT type.
type SupportedExdScreenLayouts uint32

const (
	SupportedExdScreenLayoutsFullScreen                SupportedExdScreenLayouts = 0x00000001
	SupportedExdScreenLayoutsHalfVertical              SupportedExdScreenLayouts = 0x00000002
	SupportedExdScreenLayoutsHalfHorizontal            SupportedExdScreenLayouts = 0x00000004
	SupportedExdScreenLayoutsHalfVerticalRightSplit    SupportedExdScreenLayouts = 0x00000008
	SupportedExdScreenLayoutsHalfHorizontalBottomSplit SupportedExdScreenLayouts = 0x00000010
	SupportedExdScreenLayoutsFullQuarterSplit          SupportedExdScreenLayouts = 0x00000020
	SupportedExdScreenLayoutsHalfVerticalLeftSplit     SupportedExdScreenLayouts = 0x00000040
	SupportedExdScreenLayoutsHalfHorizontalTopSplit    SupportedExdScreenLayouts = 0x00000080
	SupportedExdScreenLayoutsInvalid                   SupportedExdScreenLayouts = 0x00000000
)

// SwimStroke represents the swim_stroke FIT type.
type SwimStroke byte

const (
	SwimStrokeFreestyle    SwimStroke = 0
	SwimStrokeBackstroke   SwimStroke = 1
	SwimStrokeBreaststroke SwimStroke = 2
	SwimStrokeButterfly    SwimStroke = 3
	SwimStrokeDrill        SwimStroke = 4
	SwimStrokeMixed        SwimStroke = 5
	SwimStrokeIm           SwimStroke = 6 // IM is a mixed interval containing the same number of lengths for each of: Butterfly, Backstroke, Breaststroke, Freestyle, swam in that order.
	SwimStrokeInvalid      SwimStroke = 0xFF
)

// Switch represents the switch FIT type.
type Switch byte

const (
	SwitchOff     Switch = 0
	SwitchOn      Switch = 1
	SwitchAuto    Switch = 2
	SwitchInvalid Switch = 0xFF
)

// TapSensitivity represents the tap_sensitivity FIT type.
type TapSensitivity byte

const (
	TapSensitivityHigh    TapSensitivity = 0
	TapSensitivityMedium  TapSensitivity = 1
	TapSensitivityLow     TapSensitivity = 2
	TapSensitivityInvalid TapSensitivity = 0xFF
)

// TimeIntoDay represents the time_into_day FIT type.
type TimeIntoDay uint32

const (
	TimeIntoDayInvalid TimeIntoDay = 0xFFFFFFFF
)

// TimeMode represents the time_mode FIT type.
type TimeMode byte

const (
	TimeModeHour12            TimeMode = 0
	TimeModeHour24            TimeMode = 1 // Does not use a leading zero and has a colon
	TimeModeMilitary          TimeMode = 2 // Uses a leading zero and does not have a colon
	TimeModeHour12WithSeconds TimeMode = 3
	TimeModeHour24WithSeconds TimeMode = 4
	TimeModeUtc               TimeMode = 5
	TimeModeInvalid           TimeMode = 0xFF
)

// TimeZone represents the time_zone FIT type.
type TimeZone byte

const (
	TimeZoneAlmaty                   TimeZone = 0
	TimeZoneBangkok                  TimeZone = 1
	TimeZoneBombay                   TimeZone = 2
	TimeZoneBrasilia                 TimeZone = 3
	TimeZoneCairo                    TimeZone = 4
	TimeZoneCapeVerdeIs              TimeZone = 5
	TimeZoneDarwin                   TimeZone = 6
	TimeZoneEniwetok                 TimeZone = 7
	TimeZoneFiji                     TimeZone = 8
	TimeZoneHongKong                 TimeZone = 9
	TimeZoneIslamabad                TimeZone = 10
	TimeZoneKabul                    TimeZone = 11
	TimeZoneMagadan                  TimeZone = 12
	TimeZoneMidAtlantic              TimeZone = 13
	TimeZoneMoscow                   TimeZone = 14
	TimeZoneMuscat                   TimeZone = 15
	TimeZoneNewfoundland             TimeZone = 16
	TimeZoneSamoa                    TimeZone = 17
	TimeZoneSydney                   TimeZone = 18
	TimeZoneTehran                   TimeZone = 19
	TimeZoneTokyo                    TimeZone = 20
	TimeZoneUsAlaska                 TimeZone = 21
	TimeZoneUsAtlantic               TimeZone = 22
	TimeZoneUsCentral                TimeZone = 23
	TimeZoneUsEastern                TimeZone = 24
	TimeZoneUsHawaii                 TimeZone = 25
	TimeZoneUsMountain               TimeZone = 26
	TimeZoneUsPacific                TimeZone = 27
	TimeZoneOther                    TimeZone = 28
	TimeZoneAuckland                 TimeZone = 29
	TimeZoneKathmandu                TimeZone = 30
	TimeZoneEuropeWesternWet         TimeZone = 31
	TimeZoneEuropeCentralCet         TimeZone = 32
	TimeZoneEuropeEasternEet         TimeZone = 33
	TimeZoneJakarta                  TimeZone = 34
	TimeZonePerth                    TimeZone = 35
	TimeZoneAdelaide                 TimeZone = 36
	TimeZoneBrisbane                 TimeZone = 37
	TimeZoneTasmania                 TimeZone = 38
	TimeZoneIceland                  TimeZone = 39
	TimeZoneAmsterdam                TimeZone = 40
	TimeZoneAthens                   TimeZone = 41
	TimeZoneBarcelona                TimeZone = 42
	TimeZoneBerlin                   TimeZone = 43
	TimeZoneBrussels                 TimeZone = 44
	TimeZoneBudapest                 TimeZone = 45
	TimeZoneCopenhagen               TimeZone = 46
	TimeZoneDublin                   TimeZone = 47
	TimeZoneHelsinki                 TimeZone = 48
	TimeZoneLisbon                   TimeZone = 49
	TimeZoneLondon                   TimeZone = 50
	TimeZoneMadrid                   TimeZone = 51
	TimeZoneMunich                   TimeZone = 52
	TimeZoneOslo                     TimeZone = 53
	TimeZoneParis                    TimeZone = 54
	TimeZonePrague                   TimeZone = 55
	TimeZoneReykjavik                TimeZone = 56
	TimeZoneRome                     TimeZone = 57
	TimeZoneStockholm                TimeZone = 58
	TimeZoneVienna                   TimeZone = 59
	TimeZoneWarsaw                   TimeZone = 60
	TimeZoneZurich                   TimeZone = 61
	TimeZoneQuebec                   TimeZone = 62
	TimeZoneOntario                  TimeZone = 63
	TimeZoneManitoba                 TimeZone = 64
	TimeZoneSaskatchewan             TimeZone = 65
	TimeZoneAlberta                  TimeZone = 66
	TimeZoneBritishColumbia          TimeZone = 67
	TimeZoneBoise                    TimeZone = 68
	TimeZoneBoston                   TimeZone = 69
	TimeZoneChicago                  TimeZone = 70
	TimeZoneDallas                   TimeZone = 71
	TimeZoneDenver                   TimeZone = 72
	TimeZoneKansasCity               TimeZone = 73
	TimeZoneLasVegas                 TimeZone = 74
	TimeZoneLosAngeles               TimeZone = 75
	TimeZoneMiami                    TimeZone = 76
	TimeZoneMinneapolis              TimeZone = 77
	TimeZoneNewYork                  TimeZone = 78
	TimeZoneNewOrleans               TimeZone = 79
	TimeZonePhoenix                  TimeZone = 80
	TimeZoneSantaFe                  TimeZone = 81
	TimeZoneSeattle                  TimeZone = 82
	TimeZoneWashingtonDc             TimeZone = 83
	TimeZoneUsArizona                TimeZone = 84
	TimeZoneChita                    TimeZone = 85
	TimeZoneEkaterinburg             TimeZone = 86
	TimeZoneIrkutsk                  TimeZone = 87
	TimeZoneKaliningrad              TimeZone = 88
	TimeZoneKrasnoyarsk              TimeZone = 89
	TimeZoneNovosibirsk              TimeZone = 90
	TimeZonePetropavlovskKamchatskiy TimeZone = 91
	TimeZoneSamara                   TimeZone = 92
	TimeZoneVladivostok              TimeZone = 93
	TimeZoneMexicoCentral            TimeZone = 94
	TimeZoneMexicoMountain           TimeZone = 95
	TimeZoneMexicoPacific            TimeZone = 96
	TimeZoneCapeTown                 TimeZone = 97
	TimeZoneWinkhoek                 TimeZone = 98
	TimeZoneLagos                    TimeZone = 99
	TimeZoneRiyahd                   TimeZone = 100
	TimeZoneVenezuela                TimeZone = 101
	TimeZoneAustraliaLh              TimeZone = 102
	TimeZoneSantiago                 TimeZone = 103
	TimeZoneManual                   TimeZone = 253
	TimeZoneAutomatic                TimeZone = 254
	TimeZoneInvalid                  TimeZone = 0xFF
)

// TimerTrigger represents the timer_trigger FIT type.
type TimerTrigger byte

const (
	TimerTriggerManual           TimerTrigger = 0
	TimerTriggerAuto             TimerTrigger = 1
	TimerTriggerFitnessEquipment TimerTrigger = 2
	TimerTriggerInvalid          TimerTrigger = 0xFF
)

// TissueModelType represents the tissue_model_type FIT type.
type TissueModelType byte

const (
	TissueModelTypeZhl16c  TissueModelType = 0 // Buhlmann's decompression algorithm, version C
	TissueModelTypeInvalid TissueModelType = 0xFF
)

// Tone represents the tone FIT type.
type Tone byte

const (
	ToneOff            Tone = 0
	ToneTone           Tone = 1
	ToneVibrate        Tone = 2
	ToneToneAndVibrate Tone = 3
	ToneInvalid        Tone = 0xFF
)

// TotalBodyExerciseName represents the total_body_exercise_name FIT type.
type TotalBodyExerciseName uint16

const (
	TotalBodyExerciseNameBurpee                           TotalBodyExerciseName = 0
	TotalBodyExerciseNameWeightedBurpee                   TotalBodyExerciseName = 1
	TotalBodyExerciseNameBurpeeBoxJump                    TotalBodyExerciseName = 2
	TotalBodyExerciseNameWeightedBurpeeBoxJump            TotalBodyExerciseName = 3
	TotalBodyExerciseNameHighPullBurpee                   TotalBodyExerciseName = 4
	TotalBodyExerciseNameManMakers                        TotalBodyExerciseName = 5
	TotalBodyExerciseNameOneArmBurpee                     TotalBodyExerciseName = 6
	TotalBodyExerciseNameSquatThrusts                     TotalBodyExerciseName = 7
	TotalBodyExerciseNameWeightedSquatThrusts             TotalBodyExerciseName = 8
	TotalBodyExerciseNameSquatPlankPushUp                 TotalBodyExerciseName = 9
	TotalBodyExerciseNameWeightedSquatPlankPushUp         TotalBodyExerciseName = 10
	TotalBodyExerciseNameStandingTRotationBalance         TotalBodyExerciseName = 11
	TotalBodyExerciseNameWeightedStandingTRotationBalance TotalBodyExerciseName = 12
	TotalBodyExerciseNameInvalid                          TotalBodyExerciseName = 0xFFFF
)

// TricepsExtensionExerciseName represents the triceps_extension_exercise_name FIT type.
type TricepsExtensionExerciseName uint16

const (
	TricepsExtensionExerciseNameBenchDip                                     TricepsExtensionExerciseName = 0
	TricepsExtensionExerciseNameWeightedBenchDip                             TricepsExtensionExerciseName = 1
	TricepsExtensionExerciseNameBodyWeightDip                                TricepsExtensionExerciseName = 2
	TricepsExtensionExerciseNameCableKickback                                TricepsExtensionExerciseName = 3
	TricepsExtensionExerciseNameCableLyingTricepsExtension                   TricepsExtensionExerciseName = 4
	TricepsExtensionExerciseNameCableOverheadTricepsExtension                TricepsExtensionExerciseName = 5
	TricepsExtensionExerciseNameDumbbellKickback                             TricepsExtensionExerciseName = 6
	TricepsExtensionExerciseNameDumbbellLyingTricepsExtension                TricepsExtensionExerciseName = 7
	TricepsExtensionExerciseNameEzBarOverheadTricepsExtension                TricepsExtensionExerciseName = 8
	TricepsExtensionExerciseNameInclineDip                                   TricepsExtensionExerciseName = 9
	TricepsExtensionExerciseNameWeightedInclineDip                           TricepsExtensionExerciseName = 10
	TricepsExtensionExerciseNameInclineEzBarLyingTricepsExtension            TricepsExtensionExerciseName = 11
	TricepsExtensionExerciseNameLyingDumbbellPulloverToExtension             TricepsExtensionExerciseName = 12
	TricepsExtensionExerciseNameLyingEzBarTricepsExtension                   TricepsExtensionExerciseName = 13
	TricepsExtensionExerciseNameLyingTricepsExtensionToCloseGripBenchPress   TricepsExtensionExerciseName = 14
	TricepsExtensionExerciseNameOverheadDumbbellTricepsExtension             TricepsExtensionExerciseName = 15
	TricepsExtensionExerciseNameRecliningTricepsPress                        TricepsExtensionExerciseName = 16
	TricepsExtensionExerciseNameReverseGripPressdown                         TricepsExtensionExerciseName = 17
	TricepsExtensionExerciseNameReverseGripTricepsPressdown                  TricepsExtensionExerciseName = 18
	TricepsExtensionExerciseNameRopePressdown                                TricepsExtensionExerciseName = 19
	TricepsExtensionExerciseNameSeatedBarbellOverheadTricepsExtension        TricepsExtensionExerciseName = 20
	TricepsExtensionExerciseNameSeatedDumbbellOverheadTricepsExtension       TricepsExtensionExerciseName = 21
	TricepsExtensionExerciseNameSeatedEzBarOverheadTricepsExtension          TricepsExtensionExerciseName = 22
	TricepsExtensionExerciseNameSeatedSingleArmOverheadDumbbellExtension     TricepsExtensionExerciseName = 23
	TricepsExtensionExerciseNameSingleArmDumbbellOverheadTricepsExtension    TricepsExtensionExerciseName = 24
	TricepsExtensionExerciseNameSingleDumbbellSeatedOverheadTricepsExtension TricepsExtensionExerciseName = 25
	TricepsExtensionExerciseNameSingleLegBenchDipAndKick                     TricepsExtensionExerciseName = 26
	TricepsExtensionExerciseNameWeightedSingleLegBenchDipAndKick             TricepsExtensionExerciseName = 27
	TricepsExtensionExerciseNameSingleLegDip                                 TricepsExtensionExerciseName = 28
	TricepsExtensionExerciseNameWeightedSingleLegDip                         TricepsExtensionExerciseName = 29
	TricepsExtensionExerciseNameStaticLyingTricepsExtension                  TricepsExtensionExerciseName = 30
	TricepsExtensionExerciseNameSuspendedDip                                 TricepsExtensionExerciseName = 31
	TricepsExtensionExerciseNameWeightedSuspendedDip                         TricepsExtensionExerciseName = 32
	TricepsExtensionExerciseNameSwissBallDumbbellLyingTricepsExtension       TricepsExtensionExerciseName = 33
	TricepsExtensionExerciseNameSwissBallEzBarLyingTricepsExtension          TricepsExtensionExerciseName = 34
	TricepsExtensionExerciseNameSwissBallEzBarOverheadTricepsExtension       TricepsExtensionExerciseName = 35
	TricepsExtensionExerciseNameTabletopDip                                  TricepsExtensionExerciseName = 36
	TricepsExtensionExerciseNameWeightedTabletopDip                          TricepsExtensionExerciseName = 37
	TricepsExtensionExerciseNameTricepsExtensionOnFloor                      TricepsExtensionExerciseName = 38
	TricepsExtensionExerciseNameTricepsPressdown                             TricepsExtensionExerciseName = 39
	TricepsExtensionExerciseNameWeightedDip                                  TricepsExtensionExerciseName = 40
	TricepsExtensionExerciseNameInvalid                                      TricepsExtensionExerciseName = 0xFFFF
)

// TurnType represents the turn_type FIT type.
type TurnType byte

const (
	TurnTypeArrivingIdx             TurnType = 0
	TurnTypeArrivingLeftIdx         TurnType = 1
	TurnTypeArrivingRightIdx        TurnType = 2
	TurnTypeArrivingViaIdx          TurnType = 3
	TurnTypeArrivingViaLeftIdx      TurnType = 4
	TurnTypeArrivingViaRightIdx     TurnType = 5
	TurnTypeBearKeepLeftIdx         TurnType = 6
	TurnTypeBearKeepRightIdx        TurnType = 7
	TurnTypeContinueIdx             TurnType = 8
	TurnTypeExitLeftIdx             TurnType = 9
	TurnTypeExitRightIdx            TurnType = 10
	TurnTypeFerryIdx                TurnType = 11
	TurnTypeRoundabout45Idx         TurnType = 12
	TurnTypeRoundabout90Idx         TurnType = 13
	TurnTypeRoundabout135Idx        TurnType = 14
	TurnTypeRoundabout180Idx        TurnType = 15
	TurnTypeRoundabout225Idx        TurnType = 16
	TurnTypeRoundabout270Idx        TurnType = 17
	TurnTypeRoundabout315Idx        TurnType = 18
	TurnTypeRoundabout360Idx        TurnType = 19
	TurnTypeRoundaboutNeg45Idx      TurnType = 20
	TurnTypeRoundaboutNeg90Idx      TurnType = 21
	TurnTypeRoundaboutNeg135Idx     TurnType = 22
	TurnTypeRoundaboutNeg180Idx     TurnType = 23
	TurnTypeRoundaboutNeg225Idx     TurnType = 24
	TurnTypeRoundaboutNeg270Idx     TurnType = 25
	TurnTypeRoundaboutNeg315Idx     TurnType = 26
	TurnTypeRoundaboutNeg360Idx     TurnType = 27
	TurnTypeRoundaboutGenericIdx    TurnType = 28
	TurnTypeRoundaboutNegGenericIdx TurnType = 29
	TurnTypeSharpTurnLeftIdx        TurnType = 30
	TurnTypeSharpTurnRightIdx       TurnType = 31
	TurnTypeTurnLeftIdx             TurnType = 32
	TurnTypeTurnRightIdx            TurnType = 33
	TurnTypeUturnLeftIdx            TurnType = 34
	TurnTypeUturnRightIdx           TurnType = 35
	TurnTypeIconInvIdx              TurnType = 36
	TurnTypeIconIdxCnt              TurnType = 37
	TurnTypeInvalid                 TurnType = 0xFF
)

// UserLocalId represents the user_local_id FIT type.
type UserLocalId uint16

const (
	UserLocalIdLocalMin      UserLocalId = 0x0000
	UserLocalIdLocalMax      UserLocalId = 0x000F
	UserLocalIdStationaryMin UserLocalId = 0x0010
	UserLocalIdStationaryMax UserLocalId = 0x00FF
	UserLocalIdPortableMin   UserLocalId = 0x0100
	UserLocalIdPortableMax   UserLocalId = 0xFFFE
	UserLocalIdInvalid       UserLocalId = 0xFFFF
)

// WarmUpExerciseName represents the warm_up_exercise_name FIT type.
type WarmUpExerciseName uint16

const (
	WarmUpExerciseNameQuadrupedRocking            WarmUpExerciseName = 0
	WarmUpExerciseNameNeckTilts                   WarmUpExerciseName = 1
	WarmUpExerciseNameAnkleCircles                WarmUpExerciseName = 2
	WarmUpExerciseNameAnkleDorsiflexionWithBand   WarmUpExerciseName = 3
	WarmUpExerciseNameAnkleInternalRotation       WarmUpExerciseName = 4
	WarmUpExerciseNameArmCircles                  WarmUpExerciseName = 5
	WarmUpExerciseNameBentOverReachToSky          WarmUpExerciseName = 6
	WarmUpExerciseNameCatCamel                    WarmUpExerciseName = 7
	WarmUpExerciseNameElbowToFootLunge            WarmUpExerciseName = 8
	WarmUpExerciseNameForwardAndBackwardLegSwings WarmUpExerciseName = 9
	WarmUpExerciseNameGroiners                    WarmUpExerciseName = 10
	WarmUpExerciseNameInvertedHamstringStretch    WarmUpExerciseName = 11
	WarmUpExerciseNameLateralDuckUnder            WarmUpExerciseName = 12
	WarmUpExerciseNameNeckRotations               WarmUpExerciseName = 13
	WarmUpExerciseNameOppositeArmAndLegBalance    WarmUpExerciseName = 14
	WarmUpExerciseNameReachRollAndLift            WarmUpExerciseName = 15
	WarmUpExerciseNameScorpion                    WarmUpExerciseName = 16 // Deprecated do not use
	WarmUpExerciseNameShoulderCircles             WarmUpExerciseName = 17
	WarmUpExerciseNameSideToSideLegSwings         WarmUpExerciseName = 18
	WarmUpExerciseNameSleeperStretch              WarmUpExerciseName = 19
	WarmUpExerciseNameSlideOut                    WarmUpExerciseName = 20
	WarmUpExerciseNameSwissBallHipCrossover       WarmUpExerciseName = 21
	WarmUpExerciseNameSwissBallReachRollAndLift   WarmUpExerciseName = 22
	WarmUpExerciseNameSwissBallWindshieldWipers   WarmUpExerciseName = 23
	WarmUpExerciseNameThoracicRotation            WarmUpExerciseName = 24
	WarmUpExerciseNameWalkingHighKicks            WarmUpExerciseName = 25
	WarmUpExerciseNameWalkingHighKnees            WarmUpExerciseName = 26
	WarmUpExerciseNameWalkingKneeHugs             WarmUpExerciseName = 27
	WarmUpExerciseNameWalkingLegCradles           WarmUpExerciseName = 28
	WarmUpExerciseNameWalkout                     WarmUpExerciseName = 29
	WarmUpExerciseNameWalkoutFromPushUpPosition   WarmUpExerciseName = 30
	WarmUpExerciseNameInvalid                     WarmUpExerciseName = 0xFFFF
)

// WatchfaceMode represents the watchface_mode FIT type.
type WatchfaceMode byte

const (
	WatchfaceModeDigital   WatchfaceMode = 0
	WatchfaceModeAnalog    WatchfaceMode = 1
	WatchfaceModeConnectIq WatchfaceMode = 2
	WatchfaceModeDisabled  WatchfaceMode = 3
	WatchfaceModeInvalid   WatchfaceMode = 0xFF
)

// WaterType represents the water_type FIT type.
type WaterType byte

const (
	WaterTypeFresh   WaterType = 0
	WaterTypeSalt    WaterType = 1
	WaterTypeEn13319 WaterType = 2
	WaterTypeCustom  WaterType = 3
	WaterTypeInvalid WaterType = 0xFF
)

// WeatherReport represents the weather_report FIT type.
type WeatherReport byte

const (
	WeatherReportCurrent        WeatherReport = 0
	WeatherReportForecast       WeatherReport = 1 // Deprecated use hourly_forecast instead
	WeatherReportHourlyForecast WeatherReport = 1
	WeatherReportDailyForecast  WeatherReport = 2
	WeatherReportInvalid        WeatherReport = 0xFF
)

// WeatherSevereType represents the weather_severe_type FIT type.
type WeatherSevereType byte

const (
	WeatherSevereTypeUnspecified             WeatherSevereType = 0
	WeatherSevereTypeTornado                 WeatherSevereType = 1
	WeatherSevereTypeTsunami                 WeatherSevereType = 2
	WeatherSevereTypeHurricane               WeatherSevereType = 3
	WeatherSevereTypeExtremeWind             WeatherSevereType = 4
	WeatherSevereTypeTyphoon                 WeatherSevereType = 5
	WeatherSevereTypeInlandHurricane         WeatherSevereType = 6
	WeatherSevereTypeHurricaneForceWind      WeatherSevereType = 7
	WeatherSevereTypeWaterspout              WeatherSevereType = 8
	WeatherSevereTypeSevereThunderstorm      WeatherSevereType = 9
	WeatherSevereTypeWreckhouseWinds         WeatherSevereType = 10
	WeatherSevereTypeLesSuetesWind           WeatherSevereType = 11
	WeatherSevereTypeAvalanche               WeatherSevereType = 12
	WeatherSevereTypeFlashFlood              WeatherSevereType = 13
	WeatherSevereTypeTropicalStorm           WeatherSevereType = 14
	WeatherSevereTypeInlandTropicalStorm     WeatherSevereType = 15
	WeatherSevereTypeBlizzard                WeatherSevereType = 16
	WeatherSevereTypeIceStorm                WeatherSevereType = 17
	WeatherSevereTypeFreezingRain            WeatherSevereType = 18
	WeatherSevereTypeDebrisFlow              WeatherSevereType = 19
	WeatherSevereTypeFlashFreeze             WeatherSevereType = 20
	WeatherSevereTypeDustStorm               WeatherSevereType = 21
	WeatherSevereTypeHighWind                WeatherSevereType = 22
	WeatherSevereTypeWinterStorm             WeatherSevereType = 23
	WeatherSevereTypeHeavyFreezingSpray      WeatherSevereType = 24
	WeatherSevereTypeExtremeCold             WeatherSevereType = 25
	WeatherSevereTypeWindChill               WeatherSevereType = 26
	WeatherSevereTypeColdWave                WeatherSevereType = 27
	WeatherSevereTypeHeavySnowAlert          WeatherSevereType = 28
	WeatherSevereTypeLakeEffectBlowingSnow   WeatherSevereType = 29
	WeatherSevereTypeSnowSquall              WeatherSevereType = 30
	WeatherSevereTypeLakeEffectSnow          WeatherSevereType = 31
	WeatherSevereTypeWinterWeather           WeatherSevereType = 32
	WeatherSevereTypeSleet                   WeatherSevereType = 33
	WeatherSevereTypeSnowfall                WeatherSevereType = 34
	WeatherSevereTypeSnowAndBlowingSnow      WeatherSevereType = 35
	WeatherSevereTypeBlowingSnow             WeatherSevereType = 36
	WeatherSevereTypeSnowAlert               WeatherSevereType = 37
	WeatherSevereTypeArcticOutflow           WeatherSevereType = 38
	WeatherSevereTypeFreezingDrizzle         WeatherSevereType = 39
	WeatherSevereTypeStorm                   WeatherSevereType = 40
	WeatherSevereTypeStormSurge              WeatherSevereType = 41
	WeatherSevereTypeRainfall                WeatherSevereType = 42
	WeatherSevereTypeArealFlood              WeatherSevereType = 43
	WeatherSevereTypeCoastalFlood            WeatherSevereType = 44
	WeatherSevereTypeLakeshoreFlood          WeatherSevereType = 45
	WeatherSevereTypeExcessiveHeat           WeatherSevereType = 46
	WeatherSevereTypeHeat                    WeatherSevereType = 47
	WeatherSevereTypeWeather                 WeatherSevereType = 48
	WeatherSevereTypeHighHeatAndHumidity     WeatherSevereType = 49
	WeatherSevereTypeHumidexAndHealth        WeatherSevereType = 50
	WeatherSevereTypeHumidex                 WeatherSevereType = 51
	WeatherSevereTypeGale                    WeatherSevereType = 52
	WeatherSevereTypeFreezingSpray           WeatherSevereType = 53
	WeatherSevereTypeSpecialMarine           WeatherSevereType = 54
	WeatherSevereTypeSquall                  WeatherSevereType = 55
	WeatherSevereTypeStrongWind              WeatherSevereType = 56
	WeatherSevereTypeLakeWind                WeatherSevereType = 57
	WeatherSevereTypeMarineWeather           WeatherSevereType = 58
	WeatherSevereTypeWind                    WeatherSevereType = 59
	WeatherSevereTypeSmallCraftHazardousSeas WeatherSevereType = 60
	WeatherSevereTypeHazardousSeas           WeatherSevereType = 61
	WeatherSevereTypeSmallCraft              WeatherSevereType = 62
	WeatherSevereTypeSmallCraftWinds         WeatherSevereType = 63
	WeatherSevereTypeSmallCraftRoughBar      WeatherSevereType = 64
	WeatherSevereTypeHighWaterLevel          WeatherSevereType = 65
	WeatherSevereTypeAshfall                 WeatherSevereType = 66
	WeatherSevereTypeFreezingFog             WeatherSevereType = 67
	WeatherSevereTypeDenseFog                WeatherSevereType = 68
	WeatherSevereTypeDenseSmoke              WeatherSevereType = 69
	WeatherSevereTypeBlowingDust             WeatherSevereType = 70
	WeatherSevereTypeHardFreeze              WeatherSevereType = 71
	WeatherSevereTypeFreeze                  WeatherSevereType = 72
	WeatherSevereTypeFrost                   WeatherSevereType = 73
	WeatherSevereTypeFireWeather             WeatherSevereType = 74
	WeatherSevereTypeFlood                   WeatherSevereType = 75
	WeatherSevereTypeRipTide                 WeatherSevereType = 76
	WeatherSevereTypeHighSurf                WeatherSevereType = 77
	WeatherSevereTypeSmog                    WeatherSevereType = 78
	WeatherSevereTypeAirQuality              WeatherSevereType = 79
	WeatherSevereTypeBriskWind               WeatherSevereType = 80
	WeatherSevereTypeAirStagnation           WeatherSevereType = 81
	WeatherSevereTypeLowWater                WeatherSevereType = 82
	WeatherSevereTypeHydrological            WeatherSevereType = 83
	WeatherSevereTypeSpecialWeather          WeatherSevereType = 84
	WeatherSevereTypeInvalid                 WeatherSevereType = 0xFF
)

// WeatherSeverity represents the weather_severity FIT type.
type WeatherSeverity byte

const (
	WeatherSeverityUnknown   WeatherSeverity = 0
	WeatherSeverityWarning   WeatherSeverity = 1
	WeatherSeverityWatch     WeatherSeverity = 2
	WeatherSeverityAdvisory  WeatherSeverity = 3
	WeatherSeverityStatement WeatherSeverity = 4
	WeatherSeverityInvalid   WeatherSeverity = 0xFF
)

// WeatherStatus represents the weather_status FIT type.
type WeatherStatus byte

const (
	WeatherStatusClear                  WeatherStatus = 0
	WeatherStatusPartlyCloudy           WeatherStatus = 1
	WeatherStatusMostlyCloudy           WeatherStatus = 2
	WeatherStatusRain                   WeatherStatus = 3
	WeatherStatusSnow                   WeatherStatus = 4
	WeatherStatusWindy                  WeatherStatus = 5
	WeatherStatusThunderstorms          WeatherStatus = 6
	WeatherStatusWintryMix              WeatherStatus = 7
	WeatherStatusFog                    WeatherStatus = 8
	WeatherStatusHazy                   WeatherStatus = 11
	WeatherStatusHail                   WeatherStatus = 12
	WeatherStatusScatteredShowers       WeatherStatus = 13
	WeatherStatusScatteredThunderstorms WeatherStatus = 14
	WeatherStatusUnknownPrecipitation   WeatherStatus = 15
	WeatherStatusLightRain              WeatherStatus = 16
	WeatherStatusHeavyRain              WeatherStatus = 17
	WeatherStatusLightSnow              WeatherStatus = 18
	WeatherStatusHeavySnow              WeatherStatus = 19
	WeatherStatusLightRainSnow          WeatherStatus = 20
	WeatherStatusHeavyRainSnow          WeatherStatus = 21
	WeatherStatusCloudy                 WeatherStatus = 22
	WeatherStatusInvalid                WeatherStatus = 0xFF
)

// Weight represents the weight FIT type.
type Weight uint16

const (
	WeightCalculating Weight = 0xFFFE
	WeightInvalid     Weight = 0xFFFF
)

// WktStepDuration represents the wkt_step_duration FIT type.
type WktStepDuration byte

const (
	WktStepDurationTime                               WktStepDuration = 0
	WktStepDurationDistance                           WktStepDuration = 1
	WktStepDurationHrLessThan                         WktStepDuration = 2
	WktStepDurationHrGreaterThan                      WktStepDuration = 3
	WktStepDurationCalories                           WktStepDuration = 4
	WktStepDurationOpen                               WktStepDuration = 5
	WktStepDurationRepeatUntilStepsCmplt              WktStepDuration = 6
	WktStepDurationRepeatUntilTime                    WktStepDuration = 7
	WktStepDurationRepeatUntilDistance                WktStepDuration = 8
	WktStepDurationRepeatUntilCalories                WktStepDuration = 9
	WktStepDurationRepeatUntilHrLessThan              WktStepDuration = 10
	WktStepDurationRepeatUntilHrGreaterThan           WktStepDuration = 11
	WktStepDurationRepeatUntilPowerLessThan           WktStepDuration = 12
	WktStepDurationRepeatUntilPowerGreaterThan        WktStepDuration = 13
	WktStepDurationPowerLessThan                      WktStepDuration = 14
	WktStepDurationPowerGreaterThan                   WktStepDuration = 15
	WktStepDurationTrainingPeaksTss                   WktStepDuration = 16
	WktStepDurationRepeatUntilPowerLastLapLessThan    WktStepDuration = 17
	WktStepDurationRepeatUntilMaxPowerLastLapLessThan WktStepDuration = 18
	WktStepDurationPower3sLessThan                    WktStepDuration = 19
	WktStepDurationPower10sLessThan                   WktStepDuration = 20
	WktStepDurationPower30sLessThan                   WktStepDuration = 21
	WktStepDurationPower3sGreaterThan                 WktStepDuration = 22
	WktStepDurationPower10sGreaterThan                WktStepDuration = 23
	WktStepDurationPower30sGreaterThan                WktStepDuration = 24
	WktStepDurationPowerLapLessThan                   WktStepDuration = 25
	WktStepDurationPowerLapGreaterThan                WktStepDuration = 26
	WktStepDurationRepeatUntilTrainingPeaksTss        WktStepDuration = 27
	WktStepDurationRepetitionTime                     WktStepDuration = 28
	WktStepDurationReps                               WktStepDuration = 29
	WktStepDurationTimeOnly                           WktStepDuration = 31
	WktStepDurationInvalid                            WktStepDuration = 0xFF
)

// WktStepTarget represents the wkt_step_target FIT type.
type WktStepTarget byte

const (
	WktStepTargetSpeed        WktStepTarget = 0
	WktStepTargetHeartRate    WktStepTarget = 1
	WktStepTargetOpen         WktStepTarget = 2
	WktStepTargetCadence      WktStepTarget = 3
	WktStepTargetPower        WktStepTarget = 4
	WktStepTargetGrade        WktStepTarget = 5
	WktStepTargetResistance   WktStepTarget = 6
	WktStepTargetPower3s      WktStepTarget = 7
	WktStepTargetPower10s     WktStepTarget = 8
	WktStepTargetPower30s     WktStepTarget = 9
	WktStepTargetPowerLap     WktStepTarget = 10
	WktStepTargetSwimStroke   WktStepTarget = 11
	WktStepTargetSpeedLap     WktStepTarget = 12
	WktStepTargetHeartRateLap WktStepTarget = 13
	WktStepTargetInvalid      WktStepTarget = 0xFF
)

// WorkoutCapabilities represents the workout_capabilities FIT type.
type WorkoutCapabilities uint32

const (
	WorkoutCapabilitiesInterval         WorkoutCapabilities = 0x00000001
	WorkoutCapabilitiesCustom           WorkoutCapabilities = 0x00000002
	WorkoutCapabilitiesFitnessEquipment WorkoutCapabilities = 0x00000004
	WorkoutCapabilitiesFirstbeat        WorkoutCapabilities = 0x00000008
	WorkoutCapabilitiesNewLeaf          WorkoutCapabilities = 0x00000010
	WorkoutCapabilitiesTcx              WorkoutCapabilities = 0x00000020 // For backwards compatibility. Watch should add missing id fields then clear flag.
	WorkoutCapabilitiesSpeed            WorkoutCapabilities = 0x00000080 // Speed source required for workout step.
	WorkoutCapabilitiesHeartRate        WorkoutCapabilities = 0x00000100 // Heart rate source required for workout step.
	WorkoutCapabilitiesDistance         WorkoutCapabilities = 0x00000200 // Distance source required for workout step.
	WorkoutCapabilitiesCadence          WorkoutCapabilities = 0x00000400 // Cadence source required for workout step.
	WorkoutCapabilitiesPower            WorkoutCapabilities = 0x00000800 // Power source required for workout step.
	WorkoutCapabilitiesGrade            WorkoutCapabilities = 0x00001000 // Grade source required for workout step.
	WorkoutCapabilitiesResistance       WorkoutCapabilities = 0x00002000 // Resistance source required for workout step.
	WorkoutCapabilitiesProtected        WorkoutCapabilities = 0x00004000
	WorkoutCapabilitiesInvalid          WorkoutCapabilities = 0x00000000
)

// WorkoutEquipment represents the workout_equipment FIT type.
type WorkoutEquipment byte

const (
	WorkoutEquipmentNone          WorkoutEquipment = 0
	WorkoutEquipmentSwimFins      WorkoutEquipment = 1
	WorkoutEquipmentSwimKickboard WorkoutEquipment = 2
	WorkoutEquipmentSwimPaddles   WorkoutEquipment = 3
	WorkoutEquipmentSwimPullBuoy  WorkoutEquipment = 4
	WorkoutEquipmentSwimSnorkel   WorkoutEquipment = 5
	WorkoutEquipmentInvalid       WorkoutEquipment = 0xFF
)

// WorkoutHr represents the workout_hr FIT type.
type WorkoutHr uint32

const (
	WorkoutHrBpmOffset WorkoutHr = 100
	WorkoutHrInvalid   WorkoutHr = 0xFFFFFFFF
)

// WorkoutPower represents the workout_power FIT type.
type WorkoutPower uint32

const (
	WorkoutPowerWattsOffset WorkoutPower = 1000
	WorkoutPowerInvalid     WorkoutPower = 0xFFFFFFFF
)
