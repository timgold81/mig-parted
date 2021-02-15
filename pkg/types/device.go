/*
 * Copyright (c) 2021, NVIDIA CORPORATION.  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package types

import (
	"fmt"
	"strconv"
)

type DeviceID uint32

func NewDeviceID(device, vendor uint16) DeviceID {
	return DeviceID((uint32(device) << 16) | uint32(vendor))
}

func NewDeviceIDFromString(str string) (DeviceID, error) {
	deviceID, err := strconv.ParseInt(str, 0, 32)
	if err != nil {
		return 0, fmt.Errorf("unable to create DeviceID from string '%v': %v", str, err)
	}
	return DeviceID(deviceID), nil
}

func (d DeviceID) String() string {
	return fmt.Sprintf("0x%X", uint32(d))
}

func (d DeviceID) GetVendor() uint16 {
	return uint16(d)
}

func (d DeviceID) GetDevice() uint16 {
	return uint16(d >> 16)
}
