import unittest
import countingInversions

class CountingInversionsCase(unittest.TestCase):
    def test_countInversions(self):

        test_cases = [
            {'input': None, 'expect': 0},
            {'input': [], 'expect': 0},
            {'input': [1], 'expect': 0},
            {'input': [1,2], 'expect': 0},
            {'input': [2,1], 'expect': 1},
            {'input': [1,1], 'expect': 0},
            {'input': [1,2,3,4,5], 'expect': 0},
            {'input': [5,4,3,2,1], 'expect': 10},
            {'input': [7,5,3,3,9,6,7], 'expect': 8},
            {'input': [23, 42, 32, 64, 12, 4], 'expect': 10},
            {'input': [2,1,2,2,2,2,2], 'expect': 1},
            {'input': [10,12,9,13,8,14,7], 'expect': 12},
        ]

        for test_case in test_cases:
            actual_result = countingInversions.countInversions(test_case['input'])
            self.assertEqual(test_case['expect'], actual_result)

if __name__ == '__main__':unittest.main()