// Copyright 2017 The go-mmal Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmal

var (
	EncodingH264        = FourCC('H', '2', '6', '4')
	EncodingMVC         = FourCC('M', 'V', 'C', ' ')
	EncodingH263        = FourCC('H', '2', '6', '3')
	EncodingMP4V        = FourCC('M', 'P', '4', 'V')
	EncodingMP2V        = FourCC('M', 'P', '2', 'V')
	EncodingMP1V        = FourCC('M', 'P', '1', 'V')
	EncodingWMV3        = FourCC('W', 'M', 'V', '3')
	EncodingWMV2        = FourCC('W', 'M', 'V', '2')
	EncodingWMV1        = FourCC('W', 'M', 'V', '1')
	EncodingWVC1        = FourCC('W', 'V', 'C', '1')
	EncodingVP8         = FourCC('V', 'P', '8', ' ')
	EncodingVP7         = FourCC('V', 'P', '7', ' ')
	EncodingVP6         = FourCC('V', 'P', '6', ' ')
	EncodingTHEORA      = FourCC('T', 'H', 'E', 'O')
	EncodingSPARK       = FourCC('S', 'P', 'R', 'K')
	EncodingMJPEG       = FourCC('M', 'J', 'P', 'G')
	EncodingJPEG        = FourCC('J', 'P', 'E', 'G')
	EncodingGIF         = FourCC('G', 'I', 'F', ' ')
	EncodingPNG         = FourCC('P', 'N', 'G', ' ')
	EncodingPPM         = FourCC('P', 'P', 'M', ' ')
	EncodingTGA         = FourCC('T', 'G', 'A', ' ')
	EncodingBMP         = FourCC('B', 'M', 'P', ' ')
	EncodingI420        = FourCC('I', '4', '2', '0')
	EncodingI420_SLICE  = FourCC('S', '4', '2', '0')
	EncodingYV12        = FourCC('Y', 'V', '1', '2')
	EncodingI422        = FourCC('I', '4', '2', '2')
	EncodingI422_SLICE  = FourCC('S', '4', '2', '2')
	EncodingYUYV        = FourCC('Y', 'U', 'Y', 'V')
	EncodingYVYU        = FourCC('Y', 'V', 'Y', 'U')
	EncodingUYVY        = FourCC('U', 'Y', 'V', 'Y')
	EncodingVYUY        = FourCC('V', 'Y', 'U', 'Y')
	EncodingNV12        = FourCC('N', 'V', '1', '2')
	EncodingNV21        = FourCC('N', 'V', '2', '1')
	EncodingARGB        = FourCC('A', 'R', 'G', 'B')
	EncodingARGB_SLICE  = FourCC('a', 'r', 'g', 'b')
	EncodingRGBA        = FourCC('R', 'G', 'B', 'A')
	EncodingRGBA_SLICE  = FourCC('r', 'g', 'b', 'a')
	EncodingABGR        = FourCC('A', 'B', 'G', 'R')
	EncodingABGR_SLICE  = FourCC('a', 'b', 'g', 'r')
	EncodingBGRA        = FourCC('B', 'G', 'R', 'A')
	EncodingBGRA_SLICE  = FourCC('b', 'g', 'r', 'a')
	EncodingRGB16       = FourCC('R', 'G', 'B', '2')
	EncodingRGB16_SLICE = FourCC('r', 'g', 'b', '2')
	EncodingRGB24       = FourCC('R', 'G', 'B', '3')
	EncodingRGB24_SLICE = FourCC('r', 'g', 'b', '3')
	EncodingRGB32       = FourCC('R', 'G', 'B', '4')
	EncodingRGB32_SLICE = FourCC('r', 'g', 'b', '4')
	EncodingBGR16       = FourCC('B', 'G', 'R', '2')
	EncodingBGR16_SLICE = FourCC('b', 'g', 'r', '2')
	EncodingBGR24       = FourCC('B', 'G', 'R', '3')
	EncodingBGR24_SLICE = FourCC('b', 'g', 'r', '3')
	EncodingBGR32       = FourCC('B', 'G', 'R', '4')
	EncodingBGR32_SLICE = FourCC('b', 'g', 'r', '4')

	//Bayer formats
	//FourCC values copied from V4L2 where defined.
	//10 bit per pixel packed Bayer formats.
	EncodingBayerSBGGR10P = FourCC('p', 'B', 'A', 'A') //BGGR
	EncodingBayerSGRBG10P = FourCC('p', 'g', 'A', 'A') //GRBG
	EncodingBayerSGBRG10P = FourCC('p', 'G', 'A', 'A') //GBRG
	EncodingBayerSRGGB10P = FourCC('p', 'R', 'A', 'A') //RGGB

	//8 bit per pixel Bayer formats.
	EncodingBayerSBGGR8 = FourCC('B', 'A', '8', '1') //BGGR
	EncodingBayerSGBRG8 = FourCC('G', 'B', 'R', 'G') //GBRG
	EncodingBayerSGRBG8 = FourCC('G', 'R', 'B', 'G') //GRBG
	EncodingBayerSRGGB8 = FourCC('R', 'G', 'G', 'B') //RGGB

	//12 bit per pixel Bayer formats - not defined in V4L2, only 12bit expanded to 16.
	//Copy 10bpp packed 4CC pattern
	EncodingBayerSBGGR12P = FourCC('p', 'B', '1', '2') //BGGR
	EncodingBayerSGRBG12P = FourCC('p', 'g', '1', '2') //GRBG
	EncodingBayerSGBRG12P = FourCC('p', 'G', '1', '2') //GBRG
	EncodingBayerSRGGB12P = FourCC('p', 'R', '1', '2') //RGGB

	//16 bit per pixel Bayer formats.
	EncodingBayerSBGGR16 = FourCC('R', 'G', '1', '6') //BGGR
	EncodingBayerSGBRG16 = FourCC('G', 'B', '1', '6') //GBRG
	EncodingBayerSGRBG16 = FourCC('G', 'R', '1', '6') //GRBG
	EncodingBayerSRGGB16 = FourCC('R', 'G', '1', '6') //RGGB

	//10 bit per pixel DPCM compressed to 8bits Bayer formats.
	EncodingBayerSBGGR10DPCM8 = FourCC('b', 'B', 'A', '8') //BGGR
	EncodingBayerSGBRG10DPCM8 = FourCC('b', 'G', 'A', '8') //GBRG
	EncodingBayerSGRBG10DPCM8 = FourCC('B', 'D', '1', '0') //GRBG
	EncodingBayerSRGGB10DPCM8 = FourCC('b', 'R', 'A', '8') //RGGB

	/** SAND Video (YUVUV128) format, native format understood by VideoCore.
	 * This format is *not* opaque - if requested you will receive full frames
	 * of YUV_UV video.
	 */
	EncodingYUVUV128 = FourCC('S', 'A', 'N', 'D')

	/** VideoCore opaque image format, image handles are returned to
	 * the host but not the actual image data.
	 */
	EncodingOpaque = FourCC('O', 'P', 'Q', 'V')

	/** An EGL image handle
	 */
	EncodingEGLImage = FourCC('E', 'G', 'L', 'I')

	/* }@ */

	/** \name Pre-defined audio encodings */
	/* @{ */
	EncodingPCMUnsignedBE = FourCC('P', 'C', 'M', 'U')
	EncodingPCMUnsignedLE = FourCC('p', 'c', 'm', 'u')
	EncodingPCMSignedBE   = FourCC('P', 'C', 'M', 'S')
	EncodingPCMSignedLE   = FourCC('p', 'c', 'm', 's')
	EncodingPCMFloatBE    = FourCC('P', 'C', 'M', 'F')
	EncodingPCMFloatLE    = FourCC('p', 'c', 'm', 'f')
	/* Defines for native endianness */
	// #ifdef MMAL_IS_BIG_ENDIAN
	// 	EncodingPCMUnsigned    = EncodingPCMUnsignedBE
	// 	EncodingPCMSigned      = EncodingPCMSignedBE
	// 	EncodingPCMFloat       = EncodingPCMFloatBE
	// #else
	// 	EncodingPCMUnsigned    = EncodingPCMUnsignedLE
	// 	EncodingPCMSigned      = EncodingPCMSignedLE
	// 	EncodingPCMFloat       = EncodingPCMFloatLE
	// #endif

	EncodingMP4A        = FourCC('M', 'P', '4', 'A')
	EncodingMPGA        = FourCC('M', 'P', 'G', 'A')
	EncodingALAW        = FourCC('A', 'L', 'A', 'W')
	EncodingMULAW       = FourCC('U', 'L', 'A', 'W')
	EncodingADPCMMS     = FourCC('M', 'S', 0x0, 0x2)
	EncodingADPCMIMAMS  = FourCC('M', 'S', 0x0, 0x1)
	EncodingADPCMSWF    = FourCC('A', 'S', 'W', 'F')
	EncodingWMA1        = FourCC('W', 'M', 'A', '1')
	EncodingWMA2        = FourCC('W', 'M', 'A', '2')
	EncodingWMAP        = FourCC('W', 'M', 'A', 'P')
	EncodingWMAL        = FourCC('W', 'M', 'A', 'L')
	EncodingWMAV        = FourCC('W', 'M', 'A', 'V')
	EncodingAMRNB       = FourCC('A', 'M', 'R', 'N')
	EncodingAMRWB       = FourCC('A', 'M', 'R', 'W')
	EncodingAMRWBP      = FourCC('A', 'M', 'R', 'P')
	EncodingAC3         = FourCC('A', 'C', '3', ' ')
	EncodingEAC3        = FourCC('E', 'A', 'C', '3')
	EncodingDTS         = FourCC('D', 'T', 'S', ' ')
	EncodingMLP         = FourCC('M', 'L', 'P', ' ')
	EncodingFLAC        = FourCC('F', 'L', 'A', 'C')
	EncodingVORBIS      = FourCC('V', 'O', 'R', 'B')
	EncodingSPEEX       = FourCC('S', 'P', 'X', ' ')
	EncodingATRAC3      = FourCC('A', 'T', 'R', '3')
	EncodingATRACX      = FourCC('A', 'T', 'R', 'X')
	EncodingATRACL      = FourCC('A', 'T', 'R', 'L')
	EncodingMIDI        = FourCC('M', 'I', 'D', 'I')
	EncodingEVRC        = FourCC('E', 'V', 'R', 'C')
	EncodingNELLYMOSER  = FourCC('N', 'E', 'L', 'Y')
	EncodingQCELP       = FourCC('Q', 'C', 'E', 'L')
	EncodingMP4VDIVXDRM = FourCC('M', '4', 'V', 'D')
	/* @} */

	/* @} MmalEncodings List */

	/** \defgroup MmalEncodingVariants List of pre-defined encoding variants
	 * This defines a list of common encoding variants. This list isn't exhaustive and is only
	 * provided as a convenience to avoid clients having to use FourCC codes directly.
	 * However components are allowed to define and use their own FourCC codes. */
	/* @{ */

	/** \name Pre-defined H264 encoding variants */
	/* @{ */
	/** ISO 14496-10 Annex B byte stream format */
	EncodingVARIANTH264Default = 0
	/** ISO 14496-15 AVC stream format */
	EncodingVARIANTH264AVC1 = FourCC('A', 'V', 'C', '1')
	/** Implicitly delineated NAL units without emulation prevention */
	EncodingVARIANTH264Raw = FourCC('R', 'A', 'W', ' ')
	/* @} */

	/** \name Pre-defined MPEG4 audio encoding variants */
	/* @{ */
	/** Raw stream format */
	EncodingVariantMP4ADefault = 0
	/** ADTS stream format */
	EncodingVariantMP4AADTS = FourCC('A', 'D', 'T', 'S')
	/* @} */

	/* @} MmalEncodingVariants List */

	/** \defgroup MmalColorSpace List of pre-defined video color spaces
	 * This defines a list of common color spaces. This list isn't exhaustive and is only
	 * provided as a convenience to avoid clients having to use FourCC codes directly.
	 * However components are allowed to define and use their own FourCC codes. */
	/* @{ */

	/** Unknown color space */
	ColorSpaceUnknown = 0
	/** ITU-R BT.601-5 [SDTV] */
	ColorSpaceITURBT601 = FourCC('Y', '6', '0', '1')
	/** ITU-R BT.709-3 [HDTV] */
	ColorSpaceITURBT709 = FourCC('Y', '7', '0', '9')
	/** JPEG JFIF */
	ColorSpaceJPEGJFIF = FourCC('Y', 'J', 'F', 'I')
	/** Title 47 Code of Federal Regulations (2003) 73.682 (a) (20) */
	ColorSpaceFCC = FourCC('Y', 'F', 'C', 'C')
	/** Society of Motion Picture and Television Engineers 240M (1999) */
	ColorSpaceSMPTE240M = FourCC('Y', '2', '4', '0')
	/** ITU-R BT.470-2 System M */
	ColorSpaceBT4702M = FourCC('Y', '_', '_', 'M')
	/** ITU-R BT.470-2 System BG */
	ColorSpaceBT4702BG = FourCC('Y', '_', 'B', 'G')
	/** JPEG JFIF, but with 16..255 luma */
	ColorSpaceJFIFY16255 = FourCC('Y', 'Y', '1', '6')
)
