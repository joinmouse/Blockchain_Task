// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract StringReverser {
    // 反转字符串函数
    // 参数：input - 需要反转的字符串
    // 返回值：反转后的字符串
    function reverseString(string memory input) public pure returns (string memory) {
        // 将字符串转换为字节数组
        bytes memory strBytes = bytes(input);
        // 获取字符串长度
        uint256 length = strBytes.length;
        
        // 如果字符串为空或长度为1，直接返回原字符串
        if (length <= 1) {
            return input;
        }
        
        // 创建新的字节数组用于存储反转后的结果
        bytes memory reversed = new bytes(length);
        
        // 反转字符串
        for (uint256 i = 0; i < length; i++) {
            reversed[i] = strBytes[length - 1 - i];
        }
        
        // 将字节数组转换回字符串并返回
        return string(reversed);
    }
}

contract RomanNumeral {
    // 定义罗马数字的基本值和对应的符号
    struct RomanValue {
        uint256 value;
        string symbol;
    }
    
    // 存储罗马数字的基本值和符号
    RomanValue[] private romanValues;
    
    // 存储单个罗马数字字符到值的映射
    mapping(bytes1 => uint256) private romanCharValues;
    
    constructor() {
        // 初始化单个罗马数字字符的值
        romanCharValues['I'] = 1;
        romanCharValues['V'] = 5;
        romanCharValues['X'] = 10;
        romanCharValues['L'] = 50;
        romanCharValues['C'] = 100;
        romanCharValues['D'] = 500;
        romanCharValues['M'] = 1000;
        
        // 初始化组合罗马数字的值
        romanValues.push(RomanValue(1000, "M"));
        romanValues.push(RomanValue(900, "CM"));
        romanValues.push(RomanValue(500, "D"));
        romanValues.push(RomanValue(400, "CD"));
        romanValues.push(RomanValue(100, "C"));
        romanValues.push(RomanValue(90, "XC"));
        romanValues.push(RomanValue(50, "L"));
        romanValues.push(RomanValue(40, "XL"));
        romanValues.push(RomanValue(10, "X"));
        romanValues.push(RomanValue(9, "IX"));
        romanValues.push(RomanValue(5, "V"));
        romanValues.push(RomanValue(4, "IV"));
        romanValues.push(RomanValue(1, "I"));
    }
    
    // 将整数转换为罗马数字
    // 参数：num - 需要转换的整数（1-3999）
    // 返回值：对应的罗马数字字符串
    function intToRoman(uint256 num) public view returns (string memory) {
        require(num > 0 && num < 4000, "Number must be between 1 and 3999");
        
        string memory result = "";
        
        for (uint256 i = 0; i < romanValues.length; i++) {
            while (num >= romanValues[i].value) {
                result = string(abi.encodePacked(result, romanValues[i].symbol));
                num -= romanValues[i].value;
            }
        }
        
        return result;
    }

    // 将罗马数字转换为整数
    // 参数：s - 罗马数字字符串
    // 返回值：对应的整数值
    function romanToInt(string memory s) public view returns (uint256) {
        bytes memory romanBytes = bytes(s);
        require(romanBytes.length > 0, "Empty string");
        
        uint256 result = 0;
        uint256 prevValue = 0;
        
        for (uint256 i = 0; i < romanBytes.length; i++) {
            uint256 currentValue = romanCharValues[romanBytes[i]];
            require(currentValue > 0, "Invalid Roman numeral");
            
            if (prevValue < currentValue) {
                // 处理特殊情况（如 IV, IX, XL 等）
                result += currentValue - 2 * prevValue;
            } else {
                result += currentValue;
            }
            
            prevValue = currentValue;
        }
        
        return result;
    }
}

contract ArrayMerger {
    // 合并两个有序数组
    // 参数：
    // nums1 - 第一个有序数组
    // m - 第一个数组的有效元素个数
    // nums2 - 第二个有序数组
    // n - 第二个数组的有效元素个数
    function mergeSortedArrays(
        uint256[] memory nums1,
        uint256 m,
        uint256[] memory nums2,
        uint256 n
    ) public pure returns (uint256[] memory) {
        // 创建结果数组
        uint256[] memory result = new uint256[](m + n);
        
        // 初始化三个指针
        uint256 i = 0;  // nums1的指针
        uint256 j = 0;  // nums2的指针
        uint256 k = 0;  // result的指针
        
        // 比较两个数组的元素，将较小的放入结果数组
        while (i < m && j < n) {
            if (nums1[i] <= nums2[j]) {
                result[k] = nums1[i];
                i++;
            } else {
                result[k] = nums2[j];
                j++;
            }
            k++;
        }
        
        // 处理nums1中剩余的元素
        while (i < m) {
            result[k] = nums1[i];
            i++;
            k++;
        }
        
        // 处理nums2中剩余的元素
        while (j < n) {
            result[k] = nums2[j];
            j++;
            k++;
        }
        
        return result;
    }
}

contract BinarySearch {
    // 二分查找函数
    // 参数：
    // nums - 有序数组
    // target - 要查找的目标值
    // 返回值：目标值在数组中的索引，如果不存在则返回-1
    function binarySearch(
        uint256[] memory nums,
        uint256 target
    ) public pure returns (int256) {
        // 如果数组为空，直接返回-1
        if (nums.length == 0) {
            return -1;
        }
        
        // 初始化左右边界
        uint256 left = 0;
        uint256 right = nums.length - 1;
        
        // 二分查找
        while (left <= right) {
            // 计算中间位置（避免整数溢出的写法）
            uint256 mid = left + (right - left) / 2;
            
            // 找到目标值
            if (nums[mid] == target) {
                return int256(mid);
            }
            
            // 目标值在右半部分
            if (nums[mid] < target) {
                left = mid + 1;
            }
            // 目标值在左半部分
            else {
                right = mid - 1;
            }
        }
        
        // 未找到目标值
        return -1;
    }
}
