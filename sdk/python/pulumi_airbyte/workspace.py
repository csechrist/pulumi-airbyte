# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['WorkspaceArgs', 'Workspace']

@pulumi.input_type
class WorkspaceArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Workspace resource.
        :param pulumi.Input[str] name: Name of the workspace
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the workspace
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _WorkspaceState:
    def __init__(__self__, *,
                 data_residency: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 workspace_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Workspace resources.
        :param pulumi.Input[str] data_residency: must be one of ["auto", "us", "eu"]
        :param pulumi.Input[str] name: Name of the workspace
        """
        if data_residency is not None:
            pulumi.set(__self__, "data_residency", data_residency)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if workspace_id is not None:
            pulumi.set(__self__, "workspace_id", workspace_id)

    @property
    @pulumi.getter(name="dataResidency")
    def data_residency(self) -> Optional[pulumi.Input[str]]:
        """
        must be one of ["auto", "us", "eu"]
        """
        return pulumi.get(self, "data_residency")

    @data_residency.setter
    def data_residency(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_residency", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the workspace
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="workspaceId")
    def workspace_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "workspace_id")

    @workspace_id.setter
    def workspace_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "workspace_id", value)


class Workspace(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Workspace Resource

        ## Example Usage

        ```python
        import pulumi
        import pulumi_airbyte as airbyte

        my_workspace = airbyte.Workspace("myWorkspace")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Name of the workspace
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[WorkspaceArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Workspace Resource

        ## Example Usage

        ```python
        import pulumi
        import pulumi_airbyte as airbyte

        my_workspace = airbyte.Workspace("myWorkspace")
        ```

        :param str resource_name: The name of the resource.
        :param WorkspaceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WorkspaceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WorkspaceArgs.__new__(WorkspaceArgs)

            __props__.__dict__["name"] = name
            __props__.__dict__["data_residency"] = None
            __props__.__dict__["workspace_id"] = None
        super(Workspace, __self__).__init__(
            'airbyte:index/workspace:Workspace',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            data_residency: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            workspace_id: Optional[pulumi.Input[str]] = None) -> 'Workspace':
        """
        Get an existing Workspace resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] data_residency: must be one of ["auto", "us", "eu"]
        :param pulumi.Input[str] name: Name of the workspace
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WorkspaceState.__new__(_WorkspaceState)

        __props__.__dict__["data_residency"] = data_residency
        __props__.__dict__["name"] = name
        __props__.__dict__["workspace_id"] = workspace_id
        return Workspace(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dataResidency")
    def data_residency(self) -> pulumi.Output[str]:
        """
        must be one of ["auto", "us", "eu"]
        """
        return pulumi.get(self, "data_residency")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the workspace
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="workspaceId")
    def workspace_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "workspace_id")
