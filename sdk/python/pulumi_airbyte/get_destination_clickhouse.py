# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetDestinationClickhouseResult',
    'AwaitableGetDestinationClickhouseResult',
    'get_destination_clickhouse',
    'get_destination_clickhouse_output',
]

@pulumi.output_type
class GetDestinationClickhouseResult:
    """
    A collection of values returned by getDestinationClickhouse.
    """
    def __init__(__self__, configuration=None, destination_id=None, destination_type=None, id=None, name=None, workspace_id=None):
        if configuration and not isinstance(configuration, str):
            raise TypeError("Expected argument 'configuration' to be a str")
        pulumi.set(__self__, "configuration", configuration)
        if destination_id and not isinstance(destination_id, str):
            raise TypeError("Expected argument 'destination_id' to be a str")
        pulumi.set(__self__, "destination_id", destination_id)
        if destination_type and not isinstance(destination_type, str):
            raise TypeError("Expected argument 'destination_type' to be a str")
        pulumi.set(__self__, "destination_type", destination_type)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if workspace_id and not isinstance(workspace_id, str):
            raise TypeError("Expected argument 'workspace_id' to be a str")
        pulumi.set(__self__, "workspace_id", workspace_id)

    @property
    @pulumi.getter
    def configuration(self) -> str:
        """
        The values required to configure the destination. Parsed as JSON.
        """
        return pulumi.get(self, "configuration")

    @property
    @pulumi.getter(name="destinationId")
    def destination_id(self) -> str:
        return pulumi.get(self, "destination_id")

    @property
    @pulumi.getter(name="destinationType")
    def destination_type(self) -> str:
        return pulumi.get(self, "destination_type")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="workspaceId")
    def workspace_id(self) -> str:
        return pulumi.get(self, "workspace_id")


class AwaitableGetDestinationClickhouseResult(GetDestinationClickhouseResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDestinationClickhouseResult(
            configuration=self.configuration,
            destination_id=self.destination_id,
            destination_type=self.destination_type,
            id=self.id,
            name=self.name,
            workspace_id=self.workspace_id)


def get_destination_clickhouse(destination_id: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDestinationClickhouseResult:
    """
    DestinationClickhouse DataSource

    ## Example Usage

    ```python
    import pulumi
    import pulumi_airbyte as airbyte

    my_destination_clickhouse = airbyte.get_destination_clickhouse(destination_id="...my_destination_id...")
    ```
    """
    __args__ = dict()
    __args__['destinationId'] = destination_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('airbyte:index/getDestinationClickhouse:getDestinationClickhouse', __args__, opts=opts, typ=GetDestinationClickhouseResult).value

    return AwaitableGetDestinationClickhouseResult(
        configuration=pulumi.get(__ret__, 'configuration'),
        destination_id=pulumi.get(__ret__, 'destination_id'),
        destination_type=pulumi.get(__ret__, 'destination_type'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        workspace_id=pulumi.get(__ret__, 'workspace_id'))


@_utilities.lift_output_func(get_destination_clickhouse)
def get_destination_clickhouse_output(destination_id: Optional[pulumi.Input[str]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDestinationClickhouseResult]:
    """
    DestinationClickhouse DataSource

    ## Example Usage

    ```python
    import pulumi
    import pulumi_airbyte as airbyte

    my_destination_clickhouse = airbyte.get_destination_clickhouse(destination_id="...my_destination_id...")
    ```
    """
    ...
