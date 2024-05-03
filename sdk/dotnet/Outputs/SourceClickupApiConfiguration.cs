// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Outputs
{

    [OutputType]
    public sealed class SourceClickupApiConfiguration
    {
        /// <summary>
        /// Every ClickUp API call required authentication. This field is your personal API token. See \n\nhere\n\n.
        /// </summary>
        public readonly string ApiToken;
        /// <summary>
        /// The ID of your folder in your space. Retrieve it from the `/space/{space_id}/folder` of the ClickUp API. See \n\nhere\n\n.
        /// </summary>
        public readonly string? FolderId;
        /// <summary>
        /// Include or exclude closed tasks. By default, they are excluded. See \n\nhere\n\n. Default: false
        /// </summary>
        public readonly bool? IncludeClosedTasks;
        /// <summary>
        /// The ID of your list in your folder. Retrieve it from the `/folder/{folder_id}/list` of the ClickUp API. See \n\nhere\n\n.
        /// </summary>
        public readonly string? ListId;
        /// <summary>
        /// The ID of your space in your workspace. Retrieve it from the `/team/{team_id}/space` of the ClickUp API. See \n\nhere\n\n.
        /// </summary>
        public readonly string? SpaceId;
        /// <summary>
        /// The ID of your team in ClickUp. Retrieve it from the `/team` of the ClickUp API. See \n\nhere\n\n.
        /// </summary>
        public readonly string? TeamId;

        [OutputConstructor]
        private SourceClickupApiConfiguration(
            string apiToken,

            string? folderId,

            bool? includeClosedTasks,

            string? listId,

            string? spaceId,

            string? teamId)
        {
            ApiToken = apiToken;
            FolderId = folderId;
            IncludeClosedTasks = includeClosedTasks;
            ListId = listId;
            SpaceId = spaceId;
            TeamId = teamId;
        }
    }
}