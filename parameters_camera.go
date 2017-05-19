// Copyright 2017 The go-mmal Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmal

/*
#include <interface/mmal/mmal.h>
#include <interface/mmal/mmal_parameters_camera.h>
*/
import "C"

const (
	ParameterThumbnailConfiguration   ParameterGroupCamera = (1 << 16) + iota // Takes a @ref ParameterTHUMBNAIL_CONFIG_T */ /* 0 */
	ParameterCaptureQuality                                                   // Unused?
	ParameterRotation                                                         // Takes a @ref ParameterINT32_T
	ParameterEXIFDisable                                                      // Takes a @ref ParameterBOOLEAN_T
	ParameterEXIF                                                             // Takes a @ref ParameterEXIF_T
	ParameterAwbMode                                                          // Takes a @ref MMAL_PARAM_AWBMODE_T
	ParameterImageEffect                                                      // Takes a @ref ParameterIMAGEFX_T
	ParameterColourEffect                                                     // Takes a @ref ParameterCOLOURFX_T
	ParameterFlickerAvoid                                                     // Takes a @ref ParameterFLICKERAVOID_T
	ParameterFlash                                                            // Takes a @ref ParameterFLASH_T
	ParameterRedeye                                                           // Takes a @ref ParameterREDEYE_T
	ParameterFocus                                                            // Takes a @ref ParameterFOCUS_T
	ParameterFocalLengths                                                     // Unused?
	ParameterExposureComp                                                     // Takes a @ref ParameterINT32_T or ParameterRATIONAL_T
	ParameterZoom                                                             // Takes a @ref ParameterSCALEFACTOR_T
	ParameterMirror                                                           // Takes a @ref ParameterMIRROR_T
	ParameterCameraNum                                                        // Takes a @ref ParameterUINT32_T /* 0x10 */
	ParameterCapture                                                          // Takes a @ref ParameterBOOLEAN_T
	ParameterExposureMode                                                     // Takes a @ref ParameterEXPOSUREMODE_T
	ParameterExpMeteringMode                                                  // Takes a @ref ParameterEXPOSUREMETERINGMODE_T
	ParameterFocusStatus                                                      // Takes a @ref ParameterFOCUS_STATUS_T
	ParameterCameraConfig                                                     // Takes a @ref ParameterCAMERA_CONFIG_T
	ParameterCaptureStatus                                                    // Takes a @ref ParameterCAPTURE_STATUS_T
	ParameterFaceTrack                                                        // Takes a @ref ParameterFACE_TRACK_T
	ParameterDrawBoxFacesAndFocus                                             // Takes a @ref ParameterBOOLEAN_T
	ParameterJpegQFactor                                                      // Takes a @ref ParameterUINT32_T
	ParameterFrameRate                                                        // Takes a @ref ParameterFRAME_RATE_T
	ParameterUseStc                                                           // Takes a @ref ParameterCAMERA_STC_MODE_T
	ParameterCameraInfo                                                       // Takes a @ref ParameterCAMERA_INFO_T
	ParameterVideoStabilisation                                               // Takes a @ref ParameterBOOLEAN_T
	ParameterFaceTrackResults                                                 // Takes a @ref ParameterFACE_TRACK_RESULTS_T
	ParameterEnableRawCapture                                                 // Takes a @ref ParameterBOOLEAN_T
	ParameterDPFFile                                                          // Takes a @ref ParameterURI_T /* 0x20 */
	ParameterEnableDPFFile                                                    // Takes a @ref ParameterBOOLEAN_T
	ParameterDPFFailIsFatal                                                   // Takes a @ref ParameterBOOLEAN_T
	ParameterCaptureMode                                                      // Takes a @ref ParameterCAPTUREMODE_T
	ParameterFocusRegions                                                     // Takes a @ref ParameterFOCUS_REGIONS_T
	ParameterInputCrop                                                        // Takes a @ref ParameterINPUT_CROP_T
	ParameterSensorInformation                                                // Takes a @ref ParameterSENSOR_INFORMATION_T
	ParameterFlashSelect                                                      // Takes a @ref ParameterFLASH_SELECT_T
	ParameterFieldOfView                                                      // Takes a @ref ParameterFIELD_OF_VIEW_T
	ParameterHighDynamicRange                                                 // Takes a @ref ParameterBOOLEAN_T
	ParameterDynamicRangeCompression                                          // Takes a @ref ParameterDRC_T
	ParameterAlgorithmControl                                                 // Takes a @ref ParameterALGORITHM_CONTROL_T
	ParameterSharpness                                                        // Takes a @ref ParameterRATIONAL_T
	ParameterContrast                                                         // Takes a @ref ParameterRATIONAL_T
	ParameterBrightness                                                       // Takes a @ref ParameterRATIONAL_T
	ParameterSaturation                                                       // Takes a @ref ParameterRATIONAL_T
	ParameterISO                                                              // Takes a @ref ParameterUINT32_T /* 0x30 */
	ParameterAntishake                                                        // Takes a @ref ParameterBOOLEAN_T
	ParameterImageEffectParameters                                            // Takes a @ref ParameterIMAGEFX_PARAMETERS_T
	ParameterCameraBurstCapture                                               // Takes a @ref ParameterBOOLEAN_T
	ParameterCameraMinISO                                                     // Takes a @ref ParameterUINT32_T
	ParameterCameraUseCase                                                    // Takes a @ref ParameterCAMERA_USE_CASE_T
	ParameterCaptureStatsPass                                                 // Takes a @ref ParameterBOOLEAN_T
	ParameterCameraCustomSensorConfig                                         // Takes a @ref ParameterUINT32_T
	ParameterEnableRegisterFile                                               // Takes a @ref ParameterBOOLEAN_T
	ParameterRegisterFailIsFatal                                              // Takes a @ref ParameterBOOLEAN_T
	ParameterConfigFileRegisters                                              // Takes a @ref ParameterCONFIGFILE_T
	ParameterConfigFileChunkRegisters                                         // Takes a @ref ParameterCONFIGFILE_CHUNK_T
	ParameterJpegAttachLog                                                    // Takes a @ref ParameterBOOLEAN_T
	ParameterZeroShutterLag                                                   // Takes a @ref ParameterZEROSHUTTERLAG_T
	ParameterFpsRange                                                         // Takes a @ref ParameterFPS_RANGE_T
	ParameterCaptureExposureComp                                              // Takes a @ref ParameterINT32_T
	ParameterSwSharpenDisable                                                 // Takes a @ref ParameterBOOLEAN_T /* 0x40 */
	ParameterFlashRequired                                                    // Takes a @ref ParameterBOOLEAN_T
	ParameterSwSaturationDisable                                              // Takes a @ref ParameterBOOLEAN_T
	ParameterShutterSpeed                                                     // Takes a @ref ParameterUINT32_T
	ParameterCustomAwbGains                                                   // Takes a @ref ParameterAWB_GAINS_T
	ParameterCameraSettings                                                   // Takes a @ref ParameterCAMERA_SETTINGS_T
	ParameterPrivacyIndicator                                                 // Takes a @ref ParameterPRIVACY_INDICATOR_T
	ParameterVideoDenoise                                                     // Takes a @ref ParameterBOOLEAN_T
	ParameterStillsDenoise                                                    // Takes a @ref ParameterBOOLEAN_T
	ParameterAnnotate                                                         // Takes a @ref ParameterCAMERA_ANNOTATE_T
	ParameterStereoscopicMode                                                 // Takes a @ref ParameterSTEREOSCOPIC_MODE_T
	ParameterCameraInterface                                                  // Takes a @ref ParameterCAMERA_INTERFACE_T
	ParameterCameraClockingMode                                               // Takes a @ref ParameterCAMERA_CLOCKING_MODE_T
	ParameterCameraRXConfig                                                   // Takes a @ref ParameterCAMERA_RX_CONFIG_T
	ParameterCameraRXTiming                                                   // Takes a @ref ParameterCAMERA_RX_TIMING_T
	ParameterDPFConfig                                                        // Takes a @ref ParameterUINT32_T
	ParameterJpegRestartInterval                                              // Takes a @ref ParameterUINT32_T /* 0x50 */
	ParameterCameraISPBlockOverride                                           // Takes a @ref ParameterUINT32_T
	ParameterLensShadingOverride                                              // Takes a @ref ParameterLENS_SHADING_T
)

type ParameterThumbnailConfig struct {
	c C.MMAL_PARAMETER_THUMBNAIL_CONFIG_T
}

func (p ParameterThumbnailConfig) Enable() uint32 {
	return uint32(p.c.enable)
}

func (p ParameterThumbnailConfig) Width() uint32 {
	return uint32(p.c.width)
}

func (p ParameterThumbnailConfig) Height() uint32 {
	return uint32(p.c.height)
}

func (p ParameterThumbnailConfig) Quality() uint32 {
	return uint32(p.c.quality)
}

type ParameterEXIFType struct {
	c C.MMAL_PARAMETER_EXIF_T
}

func (p ParameterEXIFType) Keylen() uint32 {
	return uint32(p.c.keylen)
}

func (p ParameterEXIFType) ValueOffset() uint32 {
	return uint32(p.c.value_offset)
}

func (p ParameterEXIFType) Valuelen() uint32 {
	return uint32(p.c.valuelen)
}

func (p ParameterEXIFType) Data() []uint8 {
	data := make([]uint8, len(p.c.data))
	for _, d := range p.c.data[:] {
		data = append(data, uint8(d))
	}
	return data
}

type ParamExposuremode C.MMAL_PARAM_EXPOSUREMODE_T

const (
	ParamExposuremodeOff ParamExposuremode = iota
	ParamExposuremodeAuto
	ParamExposuremodeNight
	ParamExposuremodeNightpreview
	ParamExposuremodeBacklight
	ParamExposuremodeSpotlight
	ParamExposuremodeSports
	ParamExposuremodeSnow
	ParamExposuremodeBeach
	ParamExposuremodeVerylong
	ParamExposuremodeFixedfps
	ParamExposuremodeAntishake
	ParamExposuremodeFireworks
	ParamExposuremodeMAX ParamExposuremode = 0x7fffffff
)

type ParameterExposuremode struct {
	c C.MMAL_PARAMETER_EXPOSUREMODE_T
}

func (p ParameterExposuremode) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterExposuremode) Value() ParamExposuremode {
	return ParamExposuremode(p.c.value)
}

type ParamExposuremeteringmode C.MMAL_PARAM_EXPOSUREMETERINGMODE_T

const (
	ParamExposuremeteringmodeAverage ParamExposuremeteringmode = iota
	ParamExposuremeteringmodeSpot
	ParamExposuremeteringmodeBacklit
	ParamExposuremeteringmodeMatrix
	ParamExposuremeteringmodeMax ParamExposuremeteringmode = 0x7fffffff
)

type ParameterExposuremeteringmode struct {
	c C.MMAL_PARAMETER_EXPOSUREMETERINGMODE_T
}

func (p ParameterExposuremeteringmode) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterExposuremeteringmode) Value() ParamExposuremeteringmode {
	return ParamExposuremeteringmode(p.c.value)
}

type ParamAWBMode C.MMAL_PARAM_AWBMODE_T

const (
	ParamAWBModeOff ParamAWBMode = iota
	ParamAWBModeAuto
	ParamAWBModeSunlight
	ParamAWBModeCloudy
	ParamAWBModeShade
	ParamAWBModeTungsten
	ParamAWBModeFluorescent
	ParamAWBModeIncandescent
	ParamAWBModeFlash
	ParamAWBModeHorizon
	ParamAWBModeMAX ParamAWBMode = 0x7fffffff
)

type ParameterAWBMode struct {
	c C.MMAL_PARAMETER_AWBMODE_T
}

func (p ParameterAWBMode) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterAWBMode) Value() ParamAWBMode {
	return ParamAWBMode(p.c.value)
}

type ParamImagefx C.MMAL_PARAM_IMAGEFX_T

const (
	ParamImagefxNone ParamImagefx = iota
	ParamImagefxNegative
	ParamImagefxSolarize
	ParamImagefxPosterize
	ParamImagefxWhiteboard
	ParamImagefxBlackboard
	ParamImagefxSketch
	ParamImagefxDenoise
	ParamImagefxEmboss
	ParamImagefxOilpaint
	ParamImagefxHatch
	ParamImagefxGpen
	ParamImagefxPastel
	ParamImagefxWatercolour
	ParamImagefxFilm
	ParamImagefxBlur
	ParamImagefxSaturation
	ParamImagefxColourswap
	ParamImagefxWashedout
	ParamImagefxPosterise
	ParamImagefxColourpoint
	ParamImagefxColourbalance
	ParamImagefxCartoon
	ParamImagefxDeinterlaceDouble
	ParamImagefxDeinterlaceADV
	ParamImagefxDeinterlaceFast
	ParamImagefxMax ParamImagefx = 0x7fffffff
)

type ParameterImagefx struct {
	c C.MMAL_PARAMETER_IMAGEFX_T
}

func (p ParameterImagefx) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterImagefx) Value() ParamImagefx {
	return ParamImagefx(p.c.value)
}

const MaxImagefxParameters = 6

type ParameterImagefxParameters struct {
	c C.MMAL_PARAMETER_IMAGEFX_PARAMETERS_T
}

func (p ParameterImagefxParameters) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterImagefxParameters) NumEffectParams() uint32 {
	return uint32(p.c.num_effect_params)
}

func (p ParameterImagefxParameters) EffectParameters() []uint32 {
	params := make([]uint32, len(p.c.effect_parameter))
	for _, p := range p.c.effect_parameter {
		params = append(params, uint32(p))
	}
	return params
}

type ParameterColourfx struct {
	c C.MMAL_PARAMETER_COLOURFX_T
}

func (p ParameterColourfx) Enable() int32 {
	return int32(p.c.enable)
}

func (p ParameterColourfx) U() uint32 {
	return uint32(p.c.u)
}

func (p ParameterColourfx) V() uint32 {
	return uint32(p.c.v)
}

type CameraSTCMode C.MMAL_CAMERA_STC_MODE_T

const (
	ParamSTCModeOff CameraSTCMode = iota
	ParamSTCModeRaw
	ParamSTCModeCooked
	ParamSTCModeMax CameraSTCMode = 0x7fffffff
)

type ParameterCameraSTCMode struct {
	c C.MMAL_PARAMETER_CAMERA_STC_MODE_T
}

func (p ParameterCameraSTCMode) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterCameraSTCMode) Value() CameraSTCMode {
	return CameraSTCMode(p.c.value)
}

type ParamFlickeravoid C.MMAL_PARAM_FLICKERAVOID_T

const (
	ParamFlickeravoidOff ParamFlickeravoid = iota
	ParamFlickeravoidAuto
	ParamFlickeravoid50hz
	ParamFlickeravoid60hz
	ParamFlickeravoidMax ParamFlickeravoid = 0x7FFFFFFF
)

type ParameterFlickeravoid struct {
	c C.MMAL_PARAMETER_FLICKERAVOID_T
}

func (p ParameterFlickeravoid) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterFlickeravoid) Value() ParamFlickeravoid {
	return ParamFlickeravoid(p.c.value)
}

type ParamFlash C.MMAL_PARAM_FLASH_T

const (
	ParamFlashOff ParamFlash = iota
	ParamFlashAuto
	ParamFlashOn
	ParamFlashRedeye
	ParamFlashFillin
	ParamFlashTorch
	ParamFlashMax ParamFlash = 0x7FFFFFFF
)

type ParameterFlashType struct {
	c C.MMAL_PARAMETER_FLASH_T
}

func (p ParameterFlashType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterFlashType) Value() ParamFlash {
	return ParamFlash(p.c.value)
}

type ParamRedeye C.MMAL_PARAM_REDEYE_T

const (
	ParamRedeyeOff ParamRedeye = iota
	ParamRedeyeOn
	ParamRedeyeSimple
	ParamRedeyeMax ParamRedeye = 0x7FFFFFFF
)

type ParameterRedeyeType struct {
	c C.MMAL_PARAMETER_REDEYE_T
}

func (p ParameterRedeyeType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterRedeyeType) Value() ParamRedeye {
	return ParamRedeye(p.c.value)
}

type ParamFocus C.MMAL_PARAM_FOCUS_T

const (
	ParamFocusAuto ParamFocus = iota
	ParamFocusAutoNear
	ParamFocusAutoMacro
	ParamFocusCaf
	ParamFocusCafNear
	ParamFocusFixedInfinity
	ParamFocusFixedHyperfocal
	ParamFocusFixedNear
	ParamFocusFixedMacro
	ParamFocusEdof
	ParamFocusCafMacro
	ParamFocusCafFast
	ParamFocusCafNearFast
	ParamFocusCafMacroFast
	ParamFocusFixedCurrent
	ParamFocusMax ParamFocus = 0x7FFFFFFF
)

type ParameterFocusType struct {
	c C.MMAL_PARAMETER_FOCUS_T
}

func (p ParameterFocusType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterFocusType) Value() ParamFocus {
	return ParamFocus(p.c.value)
}

type ParamCaptureStatus C.MMAL_PARAM_CAPTURE_STATUS_T

const (
	ParamCaptureStatusNotCapturing ParamCaptureStatus = iota
	ParamCaptureStatusCaptureStarted
	ParamCaptureStatusCaptureEnded
	ParamCaptureStatusMax ParamCaptureStatus = 0x7FFFFFFF
)

type ParameterCaptureStatusType struct {
	c C.MMAL_PARAMETER_CAPTURE_STATUS_T
}

func (p ParameterCaptureStatusType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterCaptureStatusType) Status() ParamCaptureStatus {
	return ParamCaptureStatus(p.c.status)
}

type ParamFocusStatus C.MMAL_PARAM_FOCUS_STATUS_T

const (
	ParamFocusStatusOff ParamFocusStatus = iota
	ParamFocusStatusRequest
	ParamFocusStatusReached
	ParamFocusStatusUnableToReach
	ParamFocusStatusLost
	ParamFocusStatusCafMoving
	ParamFocusStatusCafSuccess
	ParamFocusStatusCafFailed
	ParamFocusStatusManualMoving
	ParamFocusStatusManualReached
	ParamFocusStatusCafWatching
	ParamFocusStatusCafSceneChanged
	ParamFocusStatusMax ParamFocusStatus = 0x7FFFFFFF
)

type ParameterFocusStatusType struct {
	c C.MMAL_PARAMETER_CAPTURE_STATUS_T
}

func (p ParameterFocusStatusType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterFocusStatusType) Status() ParamFocusStatus {
	return ParamFocusStatus(p.c.status)
}

type ParamFaceTrackMode C.MMAL_PARAM_FACE_TRACK_MODE_T

const (
	ParamFaceDetectNone ParamFaceTrackMode = iota // Disables face detection
	ParamFaceDetectOn                             // Enables face detection
	ParamFaceDetectMax  ParamFaceTrackMode = 0x7FFFFFFF
)

type ParameterFaceTrackType struct {
	c C.MMAL_PARAMETER_FACE_TRACK_T
}

func (p ParameterFaceTrackType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterFaceTrackType) Mode() ParamFaceTrackMode {
	return ParamFaceTrackMode(p.c.mode)
}

func (p ParameterFaceTrackType) MaxRegions() uint32 {
	return uint32(p.c.maxRegions)
}

func (p ParameterFaceTrackType) Frames() uint32 {
	return uint32(p.c.frames)
}

func (p ParameterFaceTrackType) Quality() uint32 {
	return uint32(p.c.quality)
}

type ParameterFaceTrackFace struct {
	c C.MMAL_PARAMETER_FACE_TRACK_FACE_T
}

func (p ParameterFaceTrackFace) FaceID() int32 {
	return int32(p.c.face_id)
}

func (p ParameterFaceTrackFace) Score() int32 {
	return int32(p.c.score)
}

func (p ParameterFaceTrackFace) FaceRect() Rect {
	return Rect{p.c.face_rect}
}

func (p ParameterFaceTrackFace) EyeRects() []Rect {
	rects := make([]Rect, len(p.c.eye_rect))
	for _, rect := range rects {
		rects = append(rects, rect)
	}
	return rects
}

func (p ParameterFaceTrackFace) MouthRect() Rect {
	return Rect{p.c.mouth_rect}
}

type ParameterFaceTrackResultsType struct {
	c C.MMAL_PARAMETER_FACE_TRACK_RESULTS_T
}

func (p ParameterFaceTrackResultsType) NumFaces() uint32 {
	return uint32(p.c.num_faces)
}

func (p ParameterFaceTrackResultsType) FrameWidth() uint32 {
	return uint32(p.c.frame_width)
}

func (p ParameterFaceTrackResultsType) FrameHeight() uint32 {
	return uint32(p.c.frame_height)
}

func (p ParameterFaceTrackResultsType) Faces() []ParameterFaceTrackFace {
	faces := make([]ParameterFaceTrackFace, len(p.c.faces))
	for _, face := range p.c.faces {
		faces = append(faces, ParameterFaceTrackFace{face})
	}
	return faces
}

type ParameterCameraConfigTimestampMode C.MMAL_PARAMETER_CAMERA_CONFIG_TIMESTAMP_MODE_T

const (
	ParamTimestampModeZero     ParameterCameraConfigTimestampMode = iota
	ParamTimestampModeRawSTC                                      // Use the raw STC value for the frame timestamp
	ParamTimestampModeResetSTC                                    // Use the STC timestamp but subtract the timestamp of the first frame sent to give a zero based timestamp
	ParamTimestampModeMax      ParameterCameraConfigTimestampMode = 0x7FFFFFFF
)

type ParameterCameraConfigType struct {
	c C.MMAL_PARAMETER_CAMERA_CONFIG_T
}

func (p ParameterCameraConfigType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterCameraConfigType) MaxStillsW() uint32 {
	return uint32(p.c.max_stills_w)
}

func (p ParameterCameraConfigType) MaxStillsH() uint32 {
	return uint32(p.c.max_stills_h)
}

func (p ParameterCameraConfigType) StillsYuv422() uint32 {
	return uint32(p.c.stills_yuv422)
}

func (p ParameterCameraConfigType) OneShotStills() uint32 {
	return uint32(p.c.one_shot_stills)
}

func (p ParameterCameraConfigType) MaxPreviewVidoeW() uint32 {
	return uint32(p.c.max_preview_video_w)
}

func (p ParameterCameraConfigType) MaxPreviewVideoH() uint32 {
	return uint32(p.c.max_preview_video_h)
}

func (p ParameterCameraConfigType) NumPreviewVideoFrames() uint32 {
	return uint32(p.c.num_preview_video_frames)
}

func (p ParameterCameraConfigType) StillsCaptureCircularBufferHeight() uint32 {
	return uint32(p.c.stills_capture_circular_buffer_height)
}

func (p ParameterCameraConfigType) FastPreviewResume() uint32 {
	return uint32(p.c.fast_preview_resume)
}

func (p ParameterCameraConfigType) UseSTCTimestamp() ParameterCameraConfigTimestampMode {
	return ParameterCameraConfigTimestampMode(p.c.use_stc_timestamp)
}

const (
	ParameterCameraInfoMaxCameras = 4
	ParameterCameraInfoMaxFlashes = 2
	ParameterCameraInfoMaxSTRLen  = 16
)

type ParameterCameraInfoCamera struct {
	c C.MMAL_PARAMETER_CAMERA_INFO_CAMERA_T
}

func (p ParameterCameraInfoCamera) PortID() uint32 {
	return uint32(p.c.port_id)
}

func (p ParameterCameraInfoCamera) MaxWidth() uint32 {
	return uint32(p.c.max_width)
}

func (p ParameterCameraInfoCamera) MaxHeight() uint32 {
	return uint32(p.c.max_height)
}

func (p ParameterCameraInfoCamera) LensPresent() uint32 {
	return uint32(p.c.lens_present)
}

func (p ParameterCameraInfoCamera) CameraNames() []string {
	names := make([]string, len(p.c.camera_name))
	for _, n := range p.c.camera_name {
		goN := C.GoString(&n)
		names = append(names, goN)
	}
	return names
}

type ParameterCameraInfoFlashType C.MMAL_PARAMETER_CAMERA_INFO_FLASH_TYPE_T

const (
	ParameterCameraInfoFlashTypeXenon ParameterCameraInfoFlashType = iota // Make values explicit
	ParameterCameraInfoFlashTypeLed                                       // to ensure they match
	ParameterCameraInfoFlashTypeOther                                     // values in config ini
	ParameterCameraInfoFlashTypeMax   ParameterCameraInfoFlashType = 0x7FFFFFFF
)

type ParameterCameraInfoFlash struct {
	c C.MMAL_PARAMETER_CAMERA_INFO_FLASH_T
}

func (p ParameterCameraInfoFlash) FlashType() ParameterCameraInfoFlashType {
	return ParameterCameraInfoFlashType(p.c.flash_type)
}

type ParameterCameraInfoType struct {
	c C.MMAL_PARAMETER_CAMERA_INFO_T
}

func (p ParameterCameraInfoType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterCameraInfoType) NumCameras() uint32 {
	return uint32(p.c.num_cameras)
}

func (p ParameterCameraInfoType) NumFlashes() uint32 {
	return uint32(p.c.num_flashes)
}

func (p ParameterCameraInfoType) Cameras() []ParameterCameraInfoCamera {
	cameras := make([]ParameterCameraInfoCamera, len(p.c.cameras))
	for _, camera := range p.c.cameras {
		cameras = append(cameras, ParameterCameraInfoCamera{camera})
	}
	return cameras
}

func (p ParameterCameraInfoType) Flashes() []ParameterCameraInfoFlash {
	flashes := make([]ParameterCameraInfoFlash, len(p.c.flashes))
	for _, flash := range p.c.flashes {
		flashes = append(flashes, ParameterCameraInfoFlash{flash})
	}
	return flashes
}

type ParameterCapturemodeMode C.MMAL_PARAMETER_CAPTUREMODE_MODE_T

const (
	ParamCapturemodeWaitForEnd          ParameterCapturemodeMode = iota // Resumes preview once capture is completed
	ParamCapturemodeWaitForEndAndHold                                   // Resumes preview once capture is completed, and hold the image for subsequent reprocessing
	ParamCapturemodeResumeVFImmediately                                 // Resumes preview as soon as possible once capture frame is received from the sensor
)

type ParameterCapturemode struct {
	c C.MMAL_PARAMETER_CAPTUREMODE_T
}

func (p ParameterCapturemode) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterCapturemode) Mode() ParameterCapturemodeMode {
	return ParameterCapturemodeMode(p.c.mode)
}

type ParameterFocusRegionType C.MMAL_PARAMETER_FOCUS_REGION_TYPE_T

const (
	ParameterFocusRegionTypeNormal ParameterFocusRegionType = iota
	ParameterFocusRegionTypeFace                            // Region defines a face
	ParameterFocusRegionTypeMax
)

type ParameterFocusRegion struct {
	c C.MMAL_PARAMETER_FOCUS_REGION_T
}

func (p ParameterFocusRegion) Rect() Rect {
	return Rect{p.c.rect}
}

func (p ParameterFocusRegion) Weight() uint32 {
	return uint32(p.c.weight)
}

func (p ParameterFocusRegion) Mask() uint32 {
	return uint32(p.c.mask)
}

func (p ParameterFocusRegion) Type() ParameterFocusRegionType {
	return ParameterFocusRegionType(p.c._type)
}

type ParameterFocusRegionsType struct {
	c C.MMAL_PARAMETER_FOCUS_REGIONS_T
}

func (p ParameterFocusRegionsType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterFocusRegionsType) NumRegions() uint32 {
	return uint32(p.c.num_regions)
}

func (p ParameterFocusRegionsType) LockToFaces() bool {
	return BoolT(p.c.lock_to_faces) != BoolT(0)
}

func (p ParameterFocusRegionsType) Regions() []ParameterFocusRegion {
	regions := make([]ParameterFocusRegion, len(p.c.regions))
	for _, region := range p.c.regions {
		regions = append(regions, ParameterFocusRegion{region})
	}
	return regions
}

type ParameterInputCropType struct {
	c C.MMAL_PARAMETER_INPUT_CROP_T
}

func (p ParameterInputCropType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterInputCropType) Rect() Rect {
	return Rect{p.c.rect}
}

type ParameterSensorInformationType struct {
	c C.MMAL_PARAMETER_SENSOR_INFORMATION_T
}

func (p ParameterSensorInformationType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterSensorInformationType) FNumber() Rational {
	return Rational{p.c.f_number}
}

func (p ParameterSensorInformationType) FocalLength() Rational {
	return Rational{p.c.focal_length}
}

func (p ParameterSensorInformationType) ModelID() uint32 {
	return uint32(p.c.model_id)
}

func (p ParameterSensorInformationType) ManufacturerID() uint32 {
	return uint32(p.c.manufacturer_id)
}

func (p ParameterSensorInformationType) Revision() uint32 {
	return uint32(p.c.revision)
}

type ParameterFlashSelectType struct {
	c C.MMAL_PARAMETER_FLASH_SELECT_T
}

func (p ParameterFlashSelectType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterFlashSelectType) FlashType() ParameterCameraInfoFlashType {
	return ParameterCameraInfoFlashType(p.c.flash_type)
}

type ParameterFieldOfViewType struct {
	c C.MMAL_PARAMETER_FIELD_OF_VIEW_T
}

func (p ParameterFieldOfViewType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterFieldOfViewType) FovH() Rational {
	return Rational{p.c.fov_h}
}

func (p ParameterFieldOfViewType) FovV() Rational {
	return Rational{p.c.fov_v}
}

type ParameterDRCStrength C.MMAL_PARAMETER_DRC_STRENGTH_T

const (
	ParameterDRCStrengthOff ParameterDRCStrength = iota
	ParameterDRCStrengthLow
	ParameterDRCStrengthMedium
	ParameterDRCStrengthHigh
	ParameterDRCStrengthMax ParameterDRCStrength = 0x7fffffff
)

type ParameterDRC struct {
	c C.MMAL_PARAMETER_DRC_T
}

func (p ParameterDRC) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterDRC) Strength() ParameterDRCStrength {
	return ParameterDRCStrength(p.c.strength)
}

type ParameterAlgorithmControlAlgorithms C.MMAL_PARAMETER_ALGORITHM_CONTROL_ALGORITHMS_T

const (
	ParameterAlgorithmControlAlgorithmsFacetracking ParameterAlgorithmControlAlgorithms = iota
	ParameterAlgorithmControlAlgorithmsRedeyeReduction
	ParameterAlgorithmControlAlgorithmsVideoStabilisation
	ParameterAlgorithmControlAlgorithmsWriteRaw
	ParameterAlgorithmControlAlgorithmsVideoDenoise
	ParameterAlgorithmControlAlgorithmsStillsDenoise
	ParameterAlgorithmControlAlgorithmsTemporalDenoise
	ParameterAlgorithmControlAlgorithmsAntishake
	ParameterAlgorithmControlAlgorithmsImageEffects
	ParameterAlgorithmControlAlgorithmsDynamicRangeCompression
	ParameterAlgorithmControlAlgorithmsFaceRecognition
	ParameterAlgorithmControlAlgorithmsFaceBeautification
	ParameterAlgorithmControlAlgorithmsSceneDetection
	ParameterAlgorithmControlAlgorithmsHighDynamicRange
	ParameterAlgorithmControlAlgorithmsMax ParameterAlgorithmControlAlgorithms = 0x7fffffff
)

type ParameterAlgorithmControlType struct {
	c C.MMAL_PARAMETER_ALGORITHM_CONTROL_T
}

func (p ParameterAlgorithmControlType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterAlgorithmControlType) Algorithm() ParameterAlgorithmControlAlgorithms {
	return ParameterAlgorithmControlAlgorithms(p.c.algorithm)
}

func (p ParameterAlgorithmControlType) Enabled() bool {
	return BoolT(p.c.enabled) != BoolT(0)
}

type ParamCameraUseCase C.MMAL_PARAM_CAMERA_USE_CASE_T

const (
	ParamCameraUseCaseUnknown       ParamCameraUseCase = iota // Compromise on behaviour as use case totally unknown
	ParamCameraUseCaseStillsCapture                           // Stills capture use case
	ParamCameraUseCaseVideoCapture                            // Video encode (camcorder) use case
	ParamCameraUseCaseMax           ParamCameraUseCase = 0x7fffffff
)

type ParameterCameraUseCaseType struct {
	c C.MMAL_PARAMETER_CAMERA_USE_CASE_T
}

func (p ParameterCameraUseCaseType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterCameraUseCaseType) UseCase() ParamCameraUseCase {
	return ParamCameraUseCase(p.c.use_case)
}

type ParameterFPSRange struct {
	c C.MMAL_PARAMETER_FPS_RANGE_T
}

func (p ParameterFPSRange) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterFPSRange) FPSLow() Rational {
	return Rational{p.c.fps_low}
}

func (p ParameterFPSRange) FPSHigh() Rational {
	return Rational{p.c.fps_low}
}

type ParameterZeroshutterlag struct {
	c C.MMAL_PARAMETER_ZEROSHUTTERLAG_T
}

func (p ParameterZeroshutterlag) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterZeroshutterlag) ZeroShutterLagMode() bool {
	return BoolT(p.c.zero_shutter_lag_mode) != BoolT(0)
}

func (p ParameterZeroshutterlag) ConcurrentCapture() bool {
	return BoolT(p.c.concurrent_capture) != BoolT(0)
}

type ParameterAWBGains struct {
	c C.MMAL_PARAMETER_AWB_GAINS_T
}

func (p ParameterAWBGains) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterAWBGains) RGain() Rational {
	return Rational{p.c.r_gain}
}

func (p ParameterAWBGains) BGain() Rational {
	return Rational{p.c.r_gain}
}

type ParameterCameraSettingsType struct {
	c C.MMAL_PARAMETER_CAMERA_SETTINGS_T
}

func (p ParameterCameraSettingsType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterCameraSettingsType) Exposure() uint32 {
	return uint32(p.c.exposure)
}

func (p ParameterCameraSettingsType) AnalogGain() Rational {
	return Rational{p.c.analog_gain}
}

func (p ParameterCameraSettingsType) DigitalGain() Rational {
	return Rational{p.c.digital_gain}
}

func (p ParameterCameraSettingsType) AWBRedGain() Rational {
	return Rational{p.c.awb_red_gain}
}

func (p ParameterCameraSettingsType) AWBBlueGain() Rational {
	return Rational{p.c.awb_blue_gain}
}

func (p ParameterCameraSettingsType) FucusPostion() uint32 {
	return uint32(p.c.focus_position)
}

type ParamPrivacyIndicator C.MMAL_PARAM_PRIVACY_INDICATOR_T

const (
	ParameterPrivacyIndicatorOff     ParamPrivacyIndicator = iota // Indicator will be off
	ParameterPrivacyIndicatorOn                                   // Indicator will come on just after a stills capture and and remain on for 2seconds, or will be on whilst output[1] is actively producing images.
	ParameterPrivacyIndicatorForceOn                              // Turns indicator of for 2s independent of capture status
	ParameterPrivacyIndicatorMax     ParamPrivacyIndicator = 0x7fffffff
)

type ParameterPrivacyIndicatorType struct {
	c C.MMAL_PARAMETER_PRIVACY_INDICATOR_T
}

func (p ParameterPrivacyIndicatorType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterPrivacyIndicatorType) Mode() ParamPrivacyIndicator {
	return ParamPrivacyIndicator(p.c.mode)
}

const CameraAnnotateMaxTextLen = 32

type ParameterCameraAnnotate struct {
	c C.MMAL_PARAMETER_CAMERA_ANNOTATE_T
}

func (p ParameterCameraAnnotate) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterCameraAnnotate) Enable() bool {
	return BoolT(p.c.enable) != BoolT(0)
}

func (p ParameterCameraAnnotate) Text() [CameraAnnotateMaxTextLen]string {
	texts := [CameraAnnotateMaxTextLen]string{}
	for i, text := range p.c.text {
		texts[i] = C.GoString(&text)
	}
	return texts
}

func (p ParameterCameraAnnotate) ShowShutter() bool {
	return BoolT(p.c.show_shutter) != BoolT(0)
}

func (p ParameterCameraAnnotate) ShowAnalogGain() bool {
	return BoolT(p.c.show_analog_gain) != BoolT(0)
}

func (p ParameterCameraAnnotate) ShowLens() bool {
	return BoolT(p.c.show_lens) != BoolT(0)
}

func (p ParameterCameraAnnotate) ShowCaf() bool {
	return BoolT(p.c.show_caf) != BoolT(0)
}

func (p ParameterCameraAnnotate) ShowMotion() bool {
	return BoolT(p.c.show_motion) != BoolT(0)
}

const CameraAnnotateMaxTextLenV2 = 256

type ParameterCameraAnnotateV2 struct {
	c C.MMAL_PARAMETER_CAMERA_ANNOTATE_V2_T
}

func (p ParameterCameraAnnotateV2) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterCameraAnnotateV2) Enable() bool {
	return BoolT(p.c.enable) != BoolT(0)
}

func (p ParameterCameraAnnotateV2) ShowShutter() bool {
	return BoolT(p.c.show_shutter) != BoolT(0)
}

func (p ParameterCameraAnnotateV2) ShowAnalogGain() bool {
	return BoolT(p.c.show_analog_gain) != BoolT(0)
}

func (p ParameterCameraAnnotateV2) ShowLens() bool {
	return BoolT(p.c.show_lens) != BoolT(0)
}

func (p ParameterCameraAnnotateV2) ShowCaf() bool {
	return BoolT(p.c.show_caf) != BoolT(0)
}

func (p ParameterCameraAnnotateV2) ShowMotion() bool {
	return BoolT(p.c.show_motion) != BoolT(0)
}

func (p ParameterCameraAnnotateV2) ShowFrameNum() bool {
	return BoolT(p.c.show_frame_num) != BoolT(0)
}

func (p ParameterCameraAnnotateV2) BlackTextBackground() bool {
	return BoolT(p.c.black_text_background) != BoolT(0)
}

func (p ParameterCameraAnnotateV2) Text() [CameraAnnotateMaxTextLenV2]string {
	texts := [CameraAnnotateMaxTextLenV2]string{}
	for i, text := range p.c.text {
		texts[i] = C.GoString(&text)
	}
	return texts
}

const CameraAnnotateMaxTextLenV3 = 256

type ParameterCameraAnnotateV3 struct {
	c C.MMAL_PARAMETER_CAMERA_ANNOTATE_V3_T
}

func (p ParameterCameraAnnotateV3) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterCameraAnnotateV3) Enable() bool {
	return BoolT(p.c.enable) != BoolT(0)
}

func (p ParameterCameraAnnotateV3) ShowShutter() bool {
	return BoolT(p.c.show_shutter) != BoolT(0)
}

func (p ParameterCameraAnnotateV3) ShowAnalogGain() bool {
	return BoolT(p.c.show_analog_gain) != BoolT(0)
}

func (p ParameterCameraAnnotateV3) ShowLens() bool {
	return BoolT(p.c.show_lens) != BoolT(0)
}

func (p ParameterCameraAnnotateV3) ShowCaf() bool {
	return BoolT(p.c.show_caf) != BoolT(0)
}

func (p ParameterCameraAnnotateV3) ShowMotion() bool {
	return BoolT(p.c.show_motion) != BoolT(0)
}

func (p ParameterCameraAnnotateV3) ShowFrameNum() bool {
	return BoolT(p.c.show_frame_num) != BoolT(0)
}

func (p ParameterCameraAnnotateV3) EnableTextBackground() bool {
	return BoolT(p.c.enable_text_background) != BoolT(0)
}

func (p ParameterCameraAnnotateV3) CustomBackgroundColour() bool {
	return BoolT(p.c.custom_background_colour) != BoolT(0)
}

func (p ParameterCameraAnnotateV3) CustomBackgronudY() uint8 {
	return uint8(p.c.custom_background_Y)
}

func (p ParameterCameraAnnotateV3) CustomBackgronudU() uint8 {
	return uint8(p.c.custom_background_U)
}

func (p ParameterCameraAnnotateV3) CustomBackgronudV() uint8 {
	return uint8(p.c.custom_background_V)
}

func (p ParameterCameraAnnotateV3) Dummy1() uint8 {
	return uint8(p.c.dummy1)
}

func (p ParameterCameraAnnotateV3) CustomTextColour() uint8 {
	return uint8(p.c.custom_text_colour)
}

func (p ParameterCameraAnnotateV3) CustomTextY() uint8 {
	return uint8(p.c.custom_text_Y)
}

func (p ParameterCameraAnnotateV3) CustomTextU() uint8 {
	return uint8(p.c.custom_text_U)
}

func (p ParameterCameraAnnotateV3) CustomTextV() uint8 {
	return uint8(p.c.custom_text_V)
}

func (p ParameterCameraAnnotateV3) TextSize() uint8 {
	return uint8(p.c.text_size)
}

func (p ParameterCameraAnnotateV3) Text() [CameraAnnotateMaxTextLenV3]string {
	texts := [CameraAnnotateMaxTextLenV3]string{}
	for i, text := range p.c.text {
		texts[i] = C.GoString(&text)
	}
	return texts
}

type StereoscopicMode C.MMAL_STEREOSCOPIC_MODE_T

const (
	StereoscopicModeNone       StereoscopicMode = 0
	StereoscopicModeSideBySide                  = 1
	StereoscopicModeTopBottom                   = 2
	StereoscopicModeMax        StereoscopicMode = 0x7FFFFFFF
)

type ParameterStereoscopicModeType struct {
	c C.MMAL_PARAMETER_STEREOSCOPIC_MODE_T
}

func (p ParameterStereoscopicModeType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterStereoscopicModeType) Mode() StereoscopicMode {
	return StereoscopicMode(p.c.mode)
}

func (p ParameterStereoscopicModeType) Decimate() bool {
	return BoolT(p.c.decimate) != BoolT(0)
}

func (p ParameterStereoscopicModeType) SwapEyes() bool {
	return BoolT(p.c.swap_eyes) != BoolT(0)
}

type CameraInterface C.MMAL_CAMERA_INTERFACE_T

const (
	CameraInterfaceCsi2 CameraInterface = iota
	CameraInterfaceCcp2
	CameraInterfaceCpi
	CameraInterfaceMax CameraInterface = 0x7FFFFFFF
)

type ParameterCameraInterfaceType struct {
	c C.MMAL_PARAMETER_CAMERA_INTERFACE_T
}

func (p ParameterCameraInterfaceType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterCameraInterfaceType) Mode() CameraInterface {
	return CameraInterface(p.c.mode)
}

type CameraClockingMode C.MMAL_CAMERA_CLOCKING_MODE_T

const (
	CameraClockingModeStrobe CameraClockingMode = 0
	CameraClockingModeClock                     = 1
	CameraClockingModeMax    CameraClockingMode = 0x7fffffff
)

type ParameterCameraClockingModeType struct {
	c C.MMAL_PARAMETER_CAMERA_INTERFACE_T
}

func (p ParameterCameraClockingModeType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterCameraClockingModeType) Mode() CameraInterface {
	return CameraInterface(p.c.mode)
}

type CameraRXConfigDecode C.MMAL_CAMERA_RX_CONFIG_DECODE

const (
	CameraRXConfigDecodeNone CameraRXConfigDecode = iota
	CameraRXConfigDecodeDPCM8TO10
	CameraRXConfigDecodeDPCM7TO10
	CameraRXConfigDecodeDPCM6TO10
	CameraRXConfigDecodeDPCM8TO12
	CameraRXConfigDecodeDPCM7TO12
	CameraRXConfigDecodeDPCM6TO12
	CameraRXConfigDecodeDPCM10TO14
	CameraRXConfigDecodeDPCM8TO14
	CameraRXConfigDecodeDPCM12TO16
	CameraRXConfigDecodeDPCM10TO16
	CameraRXConfigDecodeDPCM8TO16
	CameraRXConfigDecodeMax CameraRXConfigDecode = 0x7FFFFFFF
)

type CameraRXConfigEncode C.MMAL_CAMERA_RX_CONFIG_DECODE

const (
	CameraRXConfigEncodeNone CameraRXConfigEncode = iota
	CameraRXConfigEncodeDPCM10TO8
	CameraRXConfigEncodeDPCM12TO8
	CameraRXConfigEncodeDPCM14TO8
	CameraRXConfigEncodeMax CameraRXConfigEncode = 0x7FFFFFFF
)

type CameraRXConfigUnpack C.MMAL_CAMERA_RX_CONFIG_UNPACK

const (
	CameraRXConfigUnpackNone CameraRXConfigUnpack = iota
	CameraRXConfigUnpack6
	CameraRXConfigUnpack7
	CameraRXConfigUnpack8
	CameraRXConfigUnpack10
	CameraRXConfigUnpack12
	CameraRXConfigUnpack14
	CameraRXConfigUnpack16
	CameraRXConfigUnpackMax CameraRXConfigUnpack = 0x7FFFFFFF
)

type CameraRXConfigPack C.MMAL_CAMERA_RX_CONFIG_UNPACK

const (
	CameraRXConfigPackNone CameraRXConfigPack = iota
	CameraRXConfigPack8
	CameraRXConfigPack10
	CameraRXConfigPack12
	CameraRXConfigPack14
	CameraRXConfigPack16
	CameraRXConfigPackRaw10
	CameraRXConfigPackRaw12
	CameraRXConfigPackMax CameraRXConfigPack = 0x7FFFFFFF
)

type ParameterCameraRXConfigType struct {
	c C.MMAL_PARAMETER_CAMERA_RX_CONFIG_T
}

func (p ParameterCameraRXConfigType) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterCameraRXConfigType) Decode() CameraRXConfigDecode {
	return CameraRXConfigDecode(p.c.decode)
}

func (p ParameterCameraRXConfigType) Encode() CameraRXConfigEncode {
	return CameraRXConfigEncode(p.c.encode)
}

func (p ParameterCameraRXConfigType) Unpack() CameraRXConfigUnpack {
	return CameraRXConfigUnpack(p.c.unpack)
}

func (p ParameterCameraRXConfigType) Pack() CameraRXConfigPack {
	return CameraRXConfigPack(p.c.pack)
}

func (p ParameterCameraRXConfigType) DataLanes() uint32 {
	return uint32(p.c.data_lanes)
}

func (p ParameterCameraRXConfigType) EncodeBlockLength() uint32 {
	return uint32(p.c.encode_block_length)
}

func (p ParameterCameraRXConfigType) EmbeddedDataLines() uint32 {
	return uint32(p.c.embedded_data_lines)
}

func (p ParameterCameraRXConfigType) ImageID() uint32 {
	return uint32(p.c.image_id)
}

type ParameterCameraRXtiming struct {
	c C.MMAL_PARAMETER_CAMERA_RX_TIMING_T
}

func (p ParameterCameraRXtiming) Timing1() uint32 {
	return uint32(p.c.timing1)
}

func (p ParameterCameraRXtiming) Timing2() uint32 {
	return uint32(p.c.timing2)
}

func (p ParameterCameraRXtiming) Timing3() uint32 {
	return uint32(p.c.timing3)
}

func (p ParameterCameraRXtiming) Timing4() uint32 {
	return uint32(p.c.timing4)
}

func (p ParameterCameraRXtiming) Timing5() uint32 {
	return uint32(p.c.timing5)
}

func (p ParameterCameraRXtiming) Term1() uint32 {
	return uint32(p.c.term1)
}

func (p ParameterCameraRXtiming) Term2() uint32 {
	return uint32(p.c.term2)
}

func (p ParameterCameraRXtiming) CpiTiming1() uint32 {
	return uint32(p.c.cpi_timing1)
}

func (p ParameterCameraRXtiming) CpiTiming2() uint32 {
	return uint32(p.c.cpi_timing2)
}

type ParameterLensShading struct {
	c C.MMAL_PARAMETER_LENS_SHADING_T
}

func (p ParameterLensShading) Hdr() ParameterHeader {
	return ParameterHeader{p.c.hdr}
}

func (p ParameterLensShading) Enabled() bool {
	return BoolT(p.c.enabled) != BoolT(0)
}

func (p ParameterLensShading) GridCellSize() uint32 {
	return uint32(p.c.grid_cell_size)
}

func (p ParameterLensShading) GridWidth() uint32 {
	return uint32(p.c.grid_width)
}

func (p ParameterLensShading) GridStride() uint32 {
	return uint32(p.c.grid_stride)
}

func (p ParameterLensShading) GridHeight() uint32 {
	return uint32(p.c.grid_height)
}

func (p ParameterLensShading) MemHandleTable() uint32 {
	return uint32(p.c.mem_handle_table)
}

func (p ParameterLensShading) RefTransform() uint32 {
	return uint32(p.c.ref_transform)
}
