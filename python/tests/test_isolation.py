import os
import unittest

class IsolationTests(unittest.TestCase):
    def test_isolation(self):
        isolation_marker = "/tmp/cisco-datacore-takehome"

        self.assertFalse(os.path.exists(isolation_marker), "Build not properly isolated")

        with open(isolation_marker, 'wt') as f:
            pass

        self.assertTrue(os.path.exists(isolation_marker), "Unable to create build isolation marker")
