// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceLemlist
    {
        /// <summary>
        /// SourceLemlist DataSource
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Airbyte = Pulumi.Airbyte;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var mySourceLemlist = Airbyte.GetSourceLemlist.Invoke(new()
        ///     {
        ///         SourceId = "...my_source_id...",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetSourceLemlistResult> InvokeAsync(GetSourceLemlistArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceLemlistResult>("airbyte:index/getSourceLemlist:getSourceLemlist", args ?? new GetSourceLemlistArgs(), options.WithDefaults());

        /// <summary>
        /// SourceLemlist DataSource
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Airbyte = Pulumi.Airbyte;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var mySourceLemlist = Airbyte.GetSourceLemlist.Invoke(new()
        ///     {
        ///         SourceId = "...my_source_id...",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetSourceLemlistResult> Invoke(GetSourceLemlistInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceLemlistResult>("airbyte:index/getSourceLemlist:getSourceLemlist", args ?? new GetSourceLemlistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceLemlistArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceLemlistArgs()
        {
        }
        public static new GetSourceLemlistArgs Empty => new GetSourceLemlistArgs();
    }

    public sealed class GetSourceLemlistInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceLemlistInvokeArgs()
        {
        }
        public static new GetSourceLemlistInvokeArgs Empty => new GetSourceLemlistInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceLemlistResult
    {
        /// <summary>
        /// The values required to configure the source. Parsed as JSON.
        /// </summary>
        public readonly string Configuration;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string SourceId;
        public readonly string SourceType;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetSourceLemlistResult(
            string configuration,

            string id,

            string name,

            string sourceId,

            string sourceType,

            string workspaceId)
        {
            Configuration = configuration;
            Id = id;
            Name = name;
            SourceId = sourceId;
            SourceType = sourceType;
            WorkspaceId = workspaceId;
        }
    }
}