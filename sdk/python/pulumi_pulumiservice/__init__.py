# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .provider import *
from .team import *
_utilities.register(
    resource_modules="""
[
 {
  "pkg": "pulumiservice",
  "mod": "index",
  "fqn": "pulumi_pulumiservice",
  "classes": {
   "pulumiservice:index:Team": "Team"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "pulumiservice",
  "token": "pulumi:providers:pulumiservice",
  "fqn": "pulumi_pulumiservice",
  "class": "Provider"
 }
]
"""
)